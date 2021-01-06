package models

type Friend struct {
	Model

	UserId int `json:"user_id"`
	FriendId int `json:"friend_id"`

	Friend User `gorm:"foreignkey:FriendId"`
}


func GetFriends(id int) (lists []interface{} ,count int) {

	// this API like a 💩

	var friends []Friend
	db.Preload("Friend").Where("user_id=?", id).Find(&friends)

	//for friend:= range friends {
	//	lists=append(lists,friend.)
	//}

	count1 := len(friends)

	for i:=0;i<count1;i++{
		lists=append(lists,friends[i].Friend)
	}

	count = len(lists)

	return
}

func GetFriendsTotal(id int) (count int) {
	db.Model(Friend{}).Where("user_id=?",id).Count(&count)

	return
}

func IsFriend(userId,friendId int) bool {
	friend:=Friend{}
	db.Where("user_id=?",userId).Where("friend_id=?",friendId).First(&friend)
	if friend.ID < 1 {
		return false
	}

	return true
}

func AddFriend(friend Friend) Friend {
	db.Create(&friend)

	// 添加双向好友
	db.Create(&Friend{
		UserId:   friend.FriendId,
		FriendId: friend.UserId,
	})

	return friend
}

func DeleteFriend(friend Friend)  {
	db.Where(friend).Delete(friend)

	// 删除双向好友
	bothFriend:=Friend{
		UserId: friend.FriendId,
		FriendId: friend.UserId,
	}
	db.Where(bothFriend).Delete(bothFriend)
}