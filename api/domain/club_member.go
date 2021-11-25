package domain

import "github.com/Wanted-Linx/linx-backend/api/ent"

type ClubMemberRegisterRequest struct {
	ClubID int `json:"club_Id"`
}

type ClubMember struct {
	StudentID int    `json:"student_id"`
	Name      string `json:"name"`
}

type JoinedClub struct {
	ClubID       int    `json:"club_id"`
	LeaderName   string `json:"leader_name"`
	Name         string `json:"name"`
	Organization string `json:"organization"`
}

type ClubMemberService interface {
	RegisterClubMember(studentID, clubID int) error
	// GetClubMemberByClubID(clubID int) (ClubMembersDto, error)
	// GetAllClubMembers(studentID, clubID int) (ClubMembersDto, error)
}

type ClubMemberRepository interface {
	Register(studentID, clubID int) (*ent.ClubMember, error)
	// FindbyClubID(clubID int) ([]*ent.ClubMember, error)
	// FindByStudentIDAndClubID(studentID, clubID int) (*ent.ClubMember, error)
}

func ClubMembersToDto(srcClubMembers []*ent.Student) []*ClubMember {
	members := make([]*ClubMember, 0)
	for _, clubMember := range srcClubMembers {
		var member ClubMember
		member.StudentID = clubMember.ID
		member.Name = clubMember.Name
		members = append(members, &member)
	}

	return members
}

func MemberClubsToDto(srcMemberClubs []*ent.Club) []*JoinedClub {
	clubs := make([]*JoinedClub, 0)

	for _, memberClub := range srcMemberClubs {
		var club JoinedClub
		club.ClubID = memberClub.ID
		club.LeaderName = memberClub.Edges.Leader.Name
		club.Name = memberClub.Name
		club.Organization = memberClub.Organization

		clubs = append(clubs, &club)
	}

	return clubs
}
