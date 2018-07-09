package controllers

import (
	"github.com/pdmp/amai/app/models"

	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type Session struct {
	gormc.Controller
}

func (c Session) Sessions() revel.Result {
	var session []*models.Session
	c.DB.Raw(`
	SELECT Sess.Session_Date, Sess.Session_Time_Start, Sess.Session_Time_End, Sess.Session_Fee, PS.Place, Sub.Subject_Name, T.Topic_Name,
	TS.Type_Session, TP.Type_Payment
	FROM Session AS Sess, Place_Session AS PS, Subject AS Sub, Topic AS T, Type_Session AS TS, Type_Payment_session AS TP;
	`).Scan(&session)

	return c.RenderJSON(session)
}

func (c Session) MySessions() revel.Result {
	var session []*models.Session
	c.DB.Raw(`
		SELECT U.User_Email, S.Session_Date, S.Session_Time_Start, S.Session_Time_End 
		FROM public."user" AS U 
		INNER JOIN Assistance AS A
		ON U.Id = A.Id_User
		INNER JOIN Session AS S
		ON A.Id_User = S.Id;
	`).Scan(&session)
	return c.RenderJSON(session)
}

func (c Session) SessionsByTopic() revel.Result {
	var session []*models.Session
	c.DB.Raw(`
		SELECT S.Subject_Name, T.Topic_Name
		FROM Class AS C
		INNER JOIN Topic AS T
		ON C.Id = T.Id
		INNER JOIN Subject AS S
		ON T.id = S.Id;
	`).Scan(&session)
	return c.RenderJSON(session)
}

func (c Session) getSessionByPlace(id uint) revel.Result {
	var session models.Session
	c.DB.Raw("SELECT Session_Date, Session_Time_Start, Session_Time_End, Session_Fee FROM Session WHERE Id_Session_Place = ?;", id).Scan(&session)
	return c.RenderJSON(c.Response.Status)
}

func (c Session) getSessionTopicBySubject(id uint) revel.Result {
	var session models.Session
	c.DB.Raw("SELECT Session_Date, Session_Time_Start, Session_Time_End, Session_Fee FROM Session WHERE Id_Subject = ?;", id).Scan(&session)
	return c.RenderJSON(c.Response.Status)
}

func (c Session) getSessionByTopic(id uint) revel.Result {
	var session models.Session
	c.DB.Raw("SELECT Session_Date, Session_Time_Start, Session_Time_End, Session_Fee FROM Session WHERE Id_Topic = ?;", id).Scan(&session)
	return c.RenderJSON(c.Response.Status)
}

func (c Session) getSessionByType(id uint) revel.Result {
	var session models.Session
	c.DB.Raw("SELECT Session_Date, Session_Time_Start, Session_Time_End, Session_Fee FROM Session WHERE Id_Type_Session = ?;", id).Scan(&session)
	return c.RenderJSON(c.Response.Status)
}

func (c Session) getSessionByPayment(id uint) revel.Result {
	var session models.Session
	c.DB.Raw("SELECT Session_Date, Session_Time_Start, Session_Time_End, Session_Fee FROM Session WHERE Id_Type_Payment = ?;", id).Scan(&session)
	return c.RenderJSON(c.Response.Status)
}
