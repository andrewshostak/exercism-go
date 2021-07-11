package rtng

import (
	"fmt"
	"sort"
)

type RatingItem struct {
	Team    string
	Matches int
	Win     int
	Draw    int
	Loss    int
	Points  int
}

func (r *RatingItem) Update(result string) {
	r.Matches++
	if result == "win" {
		r.Win++
		r.Points += 3
	}
	if result == "draw" {
		r.Draw++
		r.Points++
	}
	if result == "loss" {
		r.Loss++
	}
}

func UpdateRating(rating []RatingItem, homeTeamName, guestTeamName, result string) []RatingItem {
	var homeTeamRating, guestTeamRating *RatingItem

	for i := range rating {
		if rating[i].Team == homeTeamName {
			homeTeamRating = &rating[i]
		}
		if rating[i].Team == guestTeamName {
			guestTeamRating = &rating[i]
		}
	}

	if homeTeamRating == nil {
		homeTeamRating = &RatingItem{Team: homeTeamName}
		rating = append(rating, *homeTeamRating)
	}
	if guestTeamRating == nil {
		guestTeamRating = &RatingItem{Team: guestTeamName}
		rating = append(rating, *guestTeamRating)
	}

	switch result {
	case "win":
		homeTeamRating.Update("win")
		guestTeamRating.Update("loss")
	case "draw":
		homeTeamRating.Update("draw")
		guestTeamRating.Update("draw")
	case "loss":
		homeTeamRating.Update("loss")
		guestTeamRating.Update("win")
	}

	for i, v := range rating {
		if v.Team == homeTeamRating.Team {
			rating[i] = *homeTeamRating
		}
		if v.Team == guestTeamRating.Team {
			rating[i] = *guestTeamRating
		}
	}

	return rating
}

func SortRating(rating []RatingItem) {
	sort.Slice(rating, func(i, j int) bool {
		if rating[i].Points != rating[j].Points {
			return rating[i].Points > rating[j].Points
		}

		if rating[i].Matches != rating[j].Matches {
			return rating[i].Matches > rating[j].Matches
		}

		return rating[i].Team < rating[j].Team
	})
}

func StringifyTable(rating []RatingItem) string {
	table := fmt.Sprintf("%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P")
	for _, item := range rating {
		table += fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n", item.Team, item.Matches, item.Win, item.Draw, item.Loss, item.Points)
	}

	return table
}
