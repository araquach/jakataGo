package controllers

import (
	"jakataGo/views"
)

func NewPage() *Page {
	return &Page{
		HomeView: views.NewView("main", "views/pages/home.gohtml"),
		ContactView: views.NewView("main", "views/pages/contact.gohtml"),
		BlogView: views.NewView("main", "views/pages/blog.gohtml"),
		BlogIndView: views.NewView("main", "views/pages/blog_ind.gohtml"),
		DetailsView: views.NewView("main", "views/pages/details.gohtml"),
		// PricesView: views.NewView("main", "views/pages/prices.gohtml"),
		RecruitmentView: views.NewView("main", "views/pages/recruitment.gohtml"),
		ReviewsView: views.NewView("main", "views/pages/reviews.gohtml"),
		SalonView: views.NewView("main", "views/pages/salon.gohtml"),
		TeamView: views.NewView("main", "views/pages/team.gohtml"),
		TeamIndView: views.NewView("main", "views/pages/teamInd.gohtml"),
	}
}

type Page struct {
	HomeView *views.View
	ContactView *views.View
	BlogView *views.View
	BlogIndView *views.View
	DetailsView *views.View
	// PricesView *views.View
	RecruitmentView *views.View
	ReviewsView *views.View
	SalonView *views.View
	TeamView *views.View
	TeamIndView *views.View
}
