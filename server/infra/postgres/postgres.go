package postgres

import (
	"context"
	"net/url"

	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/streams/vocab"
	"github.com/shima004/pactive/config"
	"github.com/shima004/pactive/domain/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) AddUser(ctx context.Context, user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUser(ctx context.Context, resource string) (vocab.ActivityStreamsPerson, error) {
	var user model.User
	// err := r.DB.First(&user, resource).Error
	err := r.DB.Where("name = ?", resource).First(&user).Error
	if err != nil {
		return nil, err
	}
	serverInfo := config.GetServerInfo()
	protocol := serverInfo.Protocol
	domain := serverInfo.Host

	person := streams.NewActivityStreamsPerson()
	id := streams.NewJSONLDIdProperty()
	id.Set(&url.URL{Scheme: protocol, Host: domain, Path: "/users/" + user.Name})
	person.SetJSONLDId(id)

	context := streams.NewActivityStreamsContextProperty()
	context.AppendIRI(&url.URL{Scheme: "https", Host: "www.w3.org", Path: "ns/activitystreams"})
	context.AppendIRI(&url.URL{Scheme: "https", Host: "w3id.org", Path: "security/v1"})
	person.SetActivityStreamsContext(context)

	url_ := streams.NewActivityStreamsUrlProperty()
	url_.AppendXMLSchemaAnyURI(&url.URL{Scheme: protocol, Host: domain, Path: "/@" + user.Name})
	person.SetActivityStreamsUrl(url_)

	inbox := streams.NewActivityStreamsInboxProperty()
	inbox.SetIRI(&url.URL{Scheme: protocol, Host: domain, Path: "/users/" + user.Name + "/inbox"})
	person.SetActivityStreamsInbox(inbox)

	outbox := streams.NewActivityStreamsOutboxProperty()
	outbox.SetIRI(&url.URL{Scheme: protocol, Host: domain, Path: "/users/" + user.Name + "/outbox"})
	person.SetActivityStreamsOutbox(outbox)

	following := streams.NewActivityStreamsFollowingProperty()
	following.SetIRI(&url.URL{Scheme: protocol, Host: domain, Path: "/users/" + user.Name + "/following"})
	person.SetActivityStreamsFollowing(following)

	followers := streams.NewActivityStreamsFollowersProperty()
	followers.SetIRI(&url.URL{Scheme: protocol, Host: domain, Path: "/users/" + user.Name + "/followers"})
	person.SetActivityStreamsFollowers(followers)

	preferredUsername := streams.NewActivityStreamsPreferredUsernameProperty()
	preferredUsername.SetXMLSchemaString(user.Name)
	person.SetActivityStreamsPreferredUsername(preferredUsername)

	name := streams.NewActivityStreamsNameProperty()
	name.AppendXMLSchemaString(user.Name)
	person.SetActivityStreamsName(name)

	summary := streams.NewActivityStreamsSummaryProperty()
	summary.AppendXMLSchemaString(user.Name)
	person.SetActivityStreamsSummary(summary)

	icon := streams.NewActivityStreamsIconProperty()
	image := streams.NewActivityStreamsImage()
	url_ = streams.NewActivityStreamsUrlProperty()
	url_.AppendIRI(&url.URL{Scheme: protocol, Host: domain, Path: "/images/default_profile.png"})
	image.SetActivityStreamsUrl(url_)
	mediaType := streams.NewActivityStreamsMediaTypeProperty()
	mediaType.Set("image/png")
	image.SetActivityStreamsMediaType(mediaType)
	icon.AppendActivityStreamsImage(image)
	person.SetActivityStreamsIcon(icon)

	// tag := streams.NewActivityStreamsTagProperty()
	// person.SetActivityStreamsTag(tag)

	// attachment := streams.NewActivityStreamsAttachmentProperty()
	// person.SetActivityStreamsAttachment(attachment)

	publicKey := streams.NewW3IDSecurityV1PublicKey()
	publicKey.SetJSONLDId(id)
	owner := streams.NewW3IDSecurityV1OwnerProperty()
	owner.SetIRI(&url.URL{Scheme: protocol, Host: domain, Path: "/users/" + user.Name})
	publicKey.SetW3IDSecurityV1Owner(owner)
	publicKeyPem := streams.NewW3IDSecurityV1PublicKeyPemProperty()
	publicKeyPem.Set(user.PublicKey)
	publicKey.SetW3IDSecurityV1PublicKeyPem(publicKeyPem)
	publickeyProperty := streams.NewW3IDSecurityV1PublicKeyProperty()
	publickeyProperty.AppendW3IDSecurityV1PublicKey(publicKey)
	person.SetW3IDSecurityV1PublicKey(publickeyProperty)

	return person, nil
}

func (r *UserRepository) GetWebFinger(ctx context.Context, resource string) (*model.WebFinger, error) {
	var user model.User
	err := r.DB.Where("name = ?", resource).First(&user).Error
	if err != nil {
		return nil, err
	}
	serverInfo := config.GetServerInfo()
	protocol := serverInfo.Protocol
	domain := serverInfo.Host

	webFinger := &model.WebFinger{
		Subject: "acct:" + user.Name + "@" + domain,
		Aliases: []string{
			protocol + "://" + domain + "/@" + user.Name,
			protocol + "://" + domain + "/users/" + user.Name,
		},
		Links: []model.Link{
			{
				Rel:  "http://webfinger.net/rel/profile-page",
				Type: "text/html",
				Href: protocol + "://" + domain + "/@" + user.Name,
			},
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: protocol + "://" + domain + "/users/" + user.Name,
			},
			{
				Rel:      "http://ostatus.org/schema/1.0/subscribe",
				Template: protocol + "://" + domain + "/authorize_interaction?uri={uri}",
			},
		},
	}

	return webFinger, nil
}
