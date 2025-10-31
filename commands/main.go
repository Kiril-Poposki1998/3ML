package commands

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

func check(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
}

type Inventory struct {
	All struct {
		Hosts map[string]interface{} `yaml:"hosts"`
	} `yaml:"all"`
}

func TerraformRun() string {
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

	// Capture the IP address from Terraform output
	cmd = exec.Command("terraform", "output", "-raw", "IP")
	output, err := cmd.Output()
	check(err)

	ipAddress := strings.TrimSpace(string(output))
	fmt.Printf("Captured IP address: %s\n", ipAddress)
	fmt.Println("Terraform apply completed successfully.")

	return ipAddress
}

func updateAnsibleInventory(ipAddress string) {
	// Path to the Ansible inventory.yaml file
	inventoryFilePath := "../../infrastructure/ansible/inventory.yaml"

	// Read the existing inventory.yaml file
	file, err := os.ReadFile(inventoryFilePath)
	check(err)

	var inventory Inventory
	err = yaml.Unmarshal(file, &inventory)
	check(err)

	// Update the hosts with the new IP address
	if inventory.All.Hosts == nil {
		inventory.All.Hosts = make(map[string]interface{})
	}
	inventory.All.Hosts["new_host"] = map[string]string{
		"ansible_host": ipAddress,
	}

	// Write the updated inventory back to the file
	updatedInventory, err := yaml.Marshal(&inventory)
	check(err)

	err = os.WriteFile(inventoryFilePath, updatedInventory, 0644)
	check(err)

	fmt.Printf("Updated inventory.yaml with IP address %s.\n", ipAddress)
}

func AnsibleRun() {
	_ = os.Chdir("../../infrastructure/ansible")

	cmd := exec.Command("ansible-playbook", "-i", "inventory.yaml", "main.yaml")
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
	ipaddr := TerraformRun()
	updateAnsibleInventory(ipaddr)
	AnsibleRun()
	syncRun()
	dockerRun()
}
