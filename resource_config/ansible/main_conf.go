package ansible

var Main = `
---
- name: Setup VM
  hosts: "{{ .host }}"

  tasks:
    {{ .DockerTasks }}
    - name: Update APT cache
      become: yes
      ansible.builtin.apt:
        update_cache: yes

    - name: Install dependencies
      become: yes
      apt:
        name:
          - nginx
          - python3-certbot-nginx

    - name: Copy nginx setting
      become: yes
      copy:
        src: ./templates/template.conf
        dest: /etc/nginx/sites-enabled/template.conf
        owner: root
        group: root
      register: temp_nginx_confg

    - name: Check nginx config
      become: yes
      command: /usr/sbin/nginx -t
      when: temp_nginx_confg.changed

    - name: Restart nginx
      become: yes
      service:
        name: nginx
        state: restarted
      when: api_conf.changed or wp_conf.changed

    - name: Renewing certificates
      become: yes
      cron: minute="0" hour="8" weekday="1" job="/usr/bin/certbot -q renew --nginx && systemctl restart nginx" state=present name=renew_certs

    {{ .DockerCronJobs }}

`

var DockerCronJobs = `- name: Delete docker images
      become: yes
      cron:
        name: "remove_images"
        minute: "0"
        hour: "8"
        weekday: "1"
        job: "docker image prune -af"
        state: present

    - name: Add cronjob for removing buildx data
      cron:
        name: "remove_buildx_data"
        minute: "0"
        hour: "8"
        weekday: "0"
        job: "docker buildx prune -af"
        state: present
`

var AnsibleDocker = `- name: Check if docker is installed
      stat:
        path: /bin/docker
      register: docker_result
    - name: Add docker repo and install
      ansible.builtin.shell: |
        apt-get update
        apt-get install ca-certificates curl
        install -m 0755 -d /etc/apt/keyrings
        curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
        chmod a+r /etc/apt/keyrings/docker.asc

        # Add the repository to Apt sources:
        echo \
          "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
          $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
          tee /etc/apt/sources.list.d/docker.list > /dev/null
        apt-get update
        apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y
      when: not docker_result.stat.exists
`
