package trigger

import "github.com/infraboard/mcenter/common/validate"

func (e *Event) Validate() error {
	return validate.Validate(e)
}

func (e *GitlabWebHookEvent) Validate() error {
	return validate.Validate(e)
}
