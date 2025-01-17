package collegeService

import (
	"bi-activity/dao/collegeDAO"
	"bi-activity/models"
	"bi-activity/models/college"
	cr "bi-activity/response/college"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type ActivityManagementDAO interface {
	GetAuditRecord(collegeId uint, status, page, size int) *cr.Result
	UpdateAuditRecord(id uint, status int)
	GetAdmissionRecord(collegeId uint, status, page, size int) *cr.Result
	UpdateAdmissionRecord(id uint, status int)
	AddActivity(c *gin.Context)
}

type ActivityManagementService struct {
	activityManagementDAO *collegeDAO.ActivityManagementDAO
}

func NewActivityManagementService(activityManagementDAO *collegeDAO.ActivityManagementDAO) *ActivityManagementService {
	return &ActivityManagementService{
		activityManagementDAO,
	}
}

func (a *ActivityManagementService) GetAuditRecord(c *gin.Context) *cr.Result {
	id, _ := c.Get("id")
	collegeId := id.(uint)
	status, _ := strconv.Atoi(c.Query("status"))
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	return a.activityManagementDAO.GetAuditRecord(collegeId, status, page, size)
}

func (a *ActivityManagementService) UpdateAuditRecord(c *gin.Context) {
	var studentActivityAudit = models.StudentActivityAudit{}
	_ = c.ShouldBindJSON(&studentActivityAudit)
	a.activityManagementDAO.UpdateAuditRecord(studentActivityAudit.ID, studentActivityAudit.Status)
}

func (a *ActivityManagementService) GetAdmissionRecord(c *gin.Context) *cr.Result {
	id, _ := c.Get("id")
	collegeId := id.(uint)
	status, _ := strconv.Atoi(c.Query("status"))
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	return a.activityManagementDAO.GetAdmissionRecord(collegeId, status, page, size)
}

func (a *ActivityManagementService) UpdateAdmissionRecord(c *gin.Context) {
	var participant = models.Participant{}
	_ = c.ShouldBindJSON(&participant)
	a.activityManagementDAO.UpdateAdmissionRecord(participant.ID, participant.Status)
}

func (a *ActivityManagementService) AddActivity(c *gin.Context) {
	// 获取请求体
	var activityRelease = college.ActivityReleaseRequest{}
	_ = c.ShouldBindJSON(&activityRelease)
	// 分离数据
	// activity
	var activity = activityRelease.GetActivity()
	id, _ := c.Get("id")
	log.Println(id)
	activityPublisherId := id.(uint)
	activity.ActivityPublisherID = activityPublisherId
	// image
	var image = activityRelease.GetImage()
	a.activityManagementDAO.AddActivity(activity, image)
}
