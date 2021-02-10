package xsltactivity

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("Setting: %s", s.ASetting)

	act := &Activity{} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}
	fmt.Println("Inside xsltactivity Eval")

	// ctx.Logger().Debugf("Input: %s", input.Xml)
	xml, xmlErr := ioutil.TempFile("", "xml")
	if xmlErr != nil {
		return false, xmlErr
	}
	defer os.Remove(xml.Name())

	fmt.Println(xml.Name())

	_, xmlWriteErr := xml.WriteString(input.Xml)
	if xmlWriteErr != nil {
		return false, xmlWriteErr
	}

	cmd := exec.Cmd{
		Args: []string{"xsltproc", input.XslFile, xml.Name()},
		Env:  os.Environ(),
		Path: "/usr/bin/xsltproc",
	}

	xmlString, cmdErr := cmd.Output()
	if cmdErr != nil {
		return false, cmdErr
	}

	// xmlString := "Boo!"

	output := &Output{OutputXml: string(xmlString)}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
