package api

import (
	"strconv"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mpaas/apps/trigger"
)

const (
	GitlabEventHeaderKey = "X-Gitlab-Event"
	GitlabEventTokenKey  = "X-Gitlab-Token"
)

// 参考文档: https://docs.gitlab.com/ee/user/project/integrations/webhook_events.html
func (h *handler) HandleGitlabEvent(r *restful.Request, w *restful.Response) {
	eventType := r.HeaderParameter(GitlabEventHeaderKey)
	serviceId := r.HeaderParameter(GitlabEventTokenKey)
	switch eventType {
	case "Push Hook":
		event := trigger.NewGitlabWebHookEvent()
		err := r.ReadEntity(event)
		if err != nil {
			response.Failed(w, err)
		}

		req := trigger.NewGitlabEvent(serviceId, event)
		req.SkipRunPipeline, err = strconv.ParseBool(r.QueryParameter("skip_run_pipeline"))
		if err != nil {
			response.Failed(w, err)
		}

		h.log.Debugf("application %s accept event: %s", serviceId, event)
		ins, err := h.svc.HandleEvent(r.Request.Context(), req)
		if err != nil {
			response.Failed(w, err)
			return
		}

		response.Success(w, ins)
	case "Tag Push Hook":
	case "Merge Request Hook":
	case "Note Hook":
	case "Issue Hook":
	}
}
