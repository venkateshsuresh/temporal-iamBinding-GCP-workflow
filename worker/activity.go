package worker

import (
	"context"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
	"os"
	"strings"
	"time"
)

func AddIAMBinding(details IamDetails) error {
	projectID := details.ProjectID
	member := fmt.Sprintf("user:%s", details.User)
	flag.Parse()

	var role string = details.Role
	ctx1 := context.TODO()
	crmService, err := cloudresourcemanager.NewService(ctx1, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	if err != nil {
		logrus.Errorf("cloudresourcemanager.NewService: %v", err)

		return err
	}

	addBinding(crmService, projectID, member, role)
	policy := getPolicy(crmService, projectID)
	var binding *cloudresourcemanager.Binding
	for _, b := range policy.Bindings {
		if b.Role == role {
			binding = b
			break
		}
	}
	fmt.Println("Role: ", binding.Role)
	fmt.Print("Members: ", strings.Join(binding.Members, ", "))

	removeMember(crmService, projectID, member, role)

	return nil
}

func addBinding(crmService *cloudresourcemanager.Service, projectID, member, role string) {

	policy := getPolicy(crmService, projectID)

	var binding *cloudresourcemanager.Binding
	for _, b := range policy.Bindings {
		if b.Role == role {
			binding = b
			break
		}
	}

	if binding != nil {
		binding.Members = append(binding.Members, member)
	} else {
		binding = &cloudresourcemanager.Binding{
			Role:    role,
			Members: []string{member},
		}
		policy.Bindings = append(policy.Bindings, binding)
	}

	setPolicy(crmService, projectID, policy)

}

func removeMember(crmService *cloudresourcemanager.Service, projectID, member, role string) {

	policy := getPolicy(crmService, projectID)

	var binding *cloudresourcemanager.Binding
	var bindingIndex int
	for i, b := range policy.Bindings {
		if b.Role == role {
			binding = b
			bindingIndex = i
			break
		}
	}

	if len(binding.Members) == 1 {
		last := len(policy.Bindings) - 1
		policy.Bindings[bindingIndex] = policy.Bindings[last]
		policy.Bindings = policy.Bindings[:last]
	} else {
		var memberIndex int
		for i, mm := range binding.Members {
			if mm == member {
				memberIndex = i
			}
		}
		last := len(policy.Bindings[bindingIndex].Members) - 1
		binding.Members[memberIndex] = binding.Members[last]
		binding.Members = binding.Members[:last]
	}

	setPolicy(crmService, projectID, policy)

}

func getPolicy(crmService *cloudresourcemanager.Service, projectID string) *cloudresourcemanager.Policy {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	request := new(cloudresourcemanager.GetIamPolicyRequest)
	policy, err := crmService.Projects.GetIamPolicy(projectID, request).Do()
	if err != nil {
		logrus.Errorf("Projects.GetIamPolicy: %v", err)
	}

	return policy
}

func setPolicy(crmService *cloudresourcemanager.Service, projectID string, policy *cloudresourcemanager.Policy) {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	request := new(cloudresourcemanager.SetIamPolicyRequest)
	request.Policy = policy
	policy, err := crmService.Projects.SetIamPolicy(projectID, request).Do()
	if err != nil {
		logrus.Errorf("Projects.SetIamPolicy: %v", err)
	}
}
