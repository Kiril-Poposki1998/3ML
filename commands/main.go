package commands

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func check(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
}

func TerraformRun() {
	// execute terraform apply command
	_ = os.Chdir("./infrastructure/terraform")

	cmd := exec.Command("terraform", "init")
	stderr, err := cmd.StderrPipe()
	check(err)

	if err := cmd.Start(); err != nil {
		check(err)
	}

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		fmt.Printf("terraform init stderr: %s\n", scanner.Text())
	}

	if err := cmd.Wait(); err != nil {
		check(err)
	}

	cmd = exec.Command("terraform", "apply", "-auto-approve")
	stderr, err = cmd.StderrPipe()
	check(err)

	if err := cmd.Start(); err != nil {
		check(err)
	}

	scanner = bufio.NewScanner(stderr)
	for scanner.Scan() {
		fmt.Printf("terraform apply stderr: %s\n", scanner.Text())
	}

	if err := cmd.Wait(); err != nil {
		check(err)
	}
	fmt.Println("Terraform apply completed successfully.")
}

func AnsibleRun() {
	// Placeholder for future implementation
	_ = os.Chdir("../../infrastructure/ansible")

	cmd := exec.Command("ansible-playbook", "-i", "hosts", "main.yaml")
	stderr, err := cmd.StderrPipe()
	check(err)

	if err := cmd.Start(); err != nil {
		check(err)
	}

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		fmt.Printf("ansible-playbook stderr: %s\n", scanner.Text())
	}

	if err := cmd.Wait(); err != nil {
		check(err)
	}
	fmt.Println("Ansible playbook completed successfully.")
}

func syncRun() {
	// Placeholder for future implementation
}

func dockerRun() {
	// Placeholder for future implementation
}

func Run() {
	TerraformRun()
	AnsibleRun()
	syncRun()
	dockerRun()
}
