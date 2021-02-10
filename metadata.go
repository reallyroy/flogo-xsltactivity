package xsltactivity

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	ASetting string `md:"aSetting,required"`
}

type Input struct {
	Xml     string `md:"xml,required"`
	XslFile string `md:"xslFile,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["xml"])
	r.Xml = strVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"xml": r.Xml,
	}
}

type Output struct {
	OutputXml string `md:"outputXml"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["outputXml"])
	o.OutputXml = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"outputXml": o.OutputXml,
	}
}
