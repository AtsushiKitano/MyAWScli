package cmd

import (
	"testing"
)

func TestgetEc2InstancesList(t *testing.T) {
	ec2Lists := getEc2InstancesList()

	if ec2Lists[0][0] != "SSH Entry" {
		t.Error("getEc2InstancesList関数のエラー")
	} else if ec2Lists[0][3] != "running" {
		t.Error("getEc2InstancesList関数のエラー")
	} else if ec2Lists[0][4] != "t2.micro" {
		t.Error("getEc2InstancesList関数のエラー")
	} else if ec2Lists[0][5] != "us-west-2" {
		t.Error("getEc2InstancesList関数のエラー")
	} else if ec2Lists[0][6] != "ebs" {
		t.Error("getEc2InstancesList関数のエラー")
	} else {
		t.Log("OK")
	}

}
