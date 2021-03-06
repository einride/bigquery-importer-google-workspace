package tables

import (
	"encoding/json"
	"strings"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	"github.com/google/uuid"
	workspace "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// UsersRow follows the structure of the WebAPI. See the official documentation for field descriptions:
// https://developers.google.com/admin-sdk/directory/reference/rest/v1/users
type UsersRow struct {
	Addresses                 []UserAddress      `bigquery:"addresses"`
	AgreedToTerms             bool               `bigquery:"agreed_to_terms"`
	Aliases                   []string           `bigquery:"aliases"`
	Archived                  bool               `bigquery:"archived"`
	ChangePasswordAtNextLogin bool               `bigquery:"change_password_at_next_login"`
	CreationTime              string             `bigquery:"creation_time"`
	CustomerID                string             `bigquery:"customer_id"`
	CustomSchemas             []CustomSchema     `bigquery:"custom_schema"`
	DeletionTime              string             `bigquery:"deletion_time"`
	Emails                    []UserEmail        `bigquery:"emails"`
	Etag                      string             `bigquery:"etag"`
	ExternalIDs               []UserExternalID   `bigquery:"external_ids"`
	Gender                    UserGender         `bigquery:"gender"`
	ID                        string             `bigquery:"id"`
	IPWhitelisted             bool               `bigquery:"ip_whitelisted"`
	IsAdmin                   bool               `bigquery:"is_admin"`
	IsDelegatedAdmin          bool               `bigquery:"is_delegated_admin"`
	IsEnforcedIn2Sv           bool               `bigquery:"is_enforced_in_2sv"`
	IsEnrolledIn2Sv           bool               `bigquery:"is_enrolled_in_2sv"`
	IsMailboxSetup            bool               `bigquery:"is_mailbox_setup"`
	Keywords                  []UserKeyword      `bigquery:"keywords"`
	Kind                      string             `bigquery:"kind"`
	Languages                 []UserLanguage     `bigquery:"languages"`
	LastLoginTime             string             `bigquery:"last_login_time"`
	Locations                 []UserLocation     `bigquery:"locations"`
	Name                      UserName           `bigquery:"name"`
	NonEditableAliases        []string           `bigquery:"non_editable_aliases"`
	Notes                     []UserAbout        `bigquery:"notes"`
	OrgUnitPath               string             `bigquery:"org_unit_path"`
	Organizations             []UserOrganization `bigquery:"organizations"`
	Phones                    []UserPhone        `bigquery:"phones"`
	PrimaryEmail              string             `bigquery:"primary_email"`
	RecoveryEmail             string             `bigquery:"recovery_email"`
	RecoveryPhone             string             `bigquery:"recovery_phone"`
	Relations                 []UserRelation     `bigquery:"relations"`
	Suspended                 bool               `bigquery:"suspended"`
	SuspensionReason          string             `bigquery:"suspension_reason"`
	ThumbnailPhotoEtag        string             `bigquery:"thumbnail_photo_etag"`
	ThumbnailPhotoURL         string             `bigquery:"thumbnail_photo_url"`
	Websites                  []UserWebsite      `bigquery:"websites"`
}

var _ Row = &UsersRow{}

type SubField interface {
	FieldType() interface{}
}

type UserAddress struct {
	Country         string `bigquery:"country"`
	CountryCode     string `bigquery:"country_code"`
	CustomType      string `bigquery:"custom_type"`
	ExtendedAddress string `bigquery:"extended_address"`
	Formatted       string `bigquery:"formatted"`
	Locality        string `bigquery:"locality"`
	PoBox           string `bigquery:"po_box"`
	PostalCode      string `bigquery:"postal_code"`
	Primary         bool   `bigquery:"primary"`
	Region          string `bigquery:"region"`
	StreetAddress   string `bigquery:"street_address"`
	Type            string `bigquery:"type"`
}

func (a *UserAddress) FieldType() interface{} {
	return UserAddress{}
}

var _ SubField = &UserAddress{}

type CustomSchema struct {
	Key   string `bigquery:"key"`
	Value string `bigquery:"value"`
}

type UserExternalID struct {
	CustomType string `json:"customType" bigquery:"custom_type"`
	Type       string `json:"type" bigquery:"type"`
	Value      string `json:"value" bigquery:"value"`
}

type UserGender struct {
	AddressMeAs  string `json:"addressMeAs" bigquery:"address_me_as"`
	CustomGender string `json:"customGender" bigquery:"custom_gender"`
	Type         string `json:"type" bigquery:"type"`
}

type UserKeyword struct {
	CustomType string `json:"customType,omitempty" bigquery:"custom_type"`
	Type       string `json:"type,omitempty" bigquery:"type"`
	Value      string `json:"value,omitempty" bigquery:"value"`
}

type UserLanguage struct {
	CustomLanguage string `json:"customLanguage,omitempty" bigquery:"custom_language"`
	LanguageCode   string `json:"languageCode,omitempty" bigquery:"language_code"`
}

type UserLocation struct {
	Area         string `json:"area,omitempty" bigquery:"area"`
	BuildingID   string `json:"buildingID,omitempty" bigquery:"building_id"`
	CustomType   string `json:"customType,omitempty" bigquery:"custom_type"`
	DeskCode     string `json:"deskCode,omitempty" bigquery:"desk_code"`
	FloorName    string `json:"floorName,omitempty" bigquery:"floor_name"`
	FloorSection string `json:"floorSection,omitempty" bigquery:"floor_selection"`
	Type         string `json:"type,omitempty" bigquery:"type"`
}

type UserName struct {
	FamilyName string `json:"familyName,omitempty" bigquery:"family_name"`
	FullName   string `json:"fullName,omitempty" bigquery:"full_name"`
	GivenName  string `json:"givenName,omitempty" bigquery:"given_name"`
}

type UserAbout struct {
	ContentType string `json:"contentType,omitempty" bigquery:"content_type"`
	Value       string `json:"value,omitempty" bigquery:"value"`
}

type UserOrganization struct {
	CostCenter         string `json:"costCenter,omitempty" bigquery:"cost_center"`
	CustomType         string `json:"customType,omitempty" bigquery:"custom_type"`
	Department         string `json:"department,omitempty" bigquery:"department"`
	Description        string `json:"description,omitempty" bigquery:"description"`
	Domain             string `json:"domain,omitempty" bigquery:"domain"`
	FullTimeEquivalent int64  `json:"fullTimeEquivalent,omitempty" bigquery:"full_time_equivalent"`
	Location           string `json:"location,omitempty" bigquery:"location"`
	Name               string `json:"name,omitempty" bigquery:"name"`
	Primary            bool   `json:"primary,omitempty" bigquery:"primary"`
	Symbol             string `json:"symbol,omitempty" bigquery:"symbol"`
	Title              string `json:"title,omitempty" bigquery:"title"`
	Type               string `json:"type,omitempty" bigquery:"type"`
}

type UserPhone struct {
	CustomType string `json:"customType,omitempty" bigquery:"custom_type"`
	Primary    bool   `json:"primary,omitempty" bigquery:"primary"`
	Type       string `json:"type,omitempty" bigquery:"type"`
	Value      string `json:"value,omitempty" bigquery:"value"`
}

type UserEmail struct {
	Address    string `bigquery:"address"`
	CustomType string `bigquery:"customType"`
	Primary    bool   `bigquery:"primary"`
	Type       string `bigquery:"type"`
}

type UserRelation struct {
	CustomType string `json:"customType,omitempty" bigquery:"custom_type"`
	Type       string `json:"type,omitempty" bigquery:"type"`
	Value      string `json:"value,omitempty" bigquery:"value"`
}

type UserWebsite struct {
	CustomType string `json:"customType,omitempty" bigquery:"custom_type"`
	Primary    bool   `json:"primary,omitempty" bigquery:"primary"`
	Type       string `json:"type,omitempty" bigquery:"type"`
	Value      string `json:"value,omitempty" bigquery:"value"`
}

func (u *UsersRow) TableID(date civil.Date) string {
	return "users_" + strings.ReplaceAll(date.String(), "-", "")
}

func (u *UsersRow) ValueSaver(jobID uuid.UUID) bigquery.ValueSaver {
	return &bigquery.StructSaver{
		Schema:   u.Schema(),
		InsertID: u.InsertID(jobID),
		Struct:   u,
	}
}

func (u *UsersRow) Schema() bigquery.Schema {
	schema, _ := bigquery.InferSchema(u)
	return schema
}

func (u *UsersRow) TableMetadata() *bigquery.TableMetadata {
	return &bigquery.TableMetadata{
		Description: "Users follows the structure of the WebAPI. See the official documentation for field descriptions:" +
			"https://developers.google.com/admin-sdk/directory/reference/rest/v1/users",
		Schema: u.Schema(),
	}
}

func (u *UsersRow) InsertID(jobID uuid.UUID) string {
	return strings.Join([]string{
		jobID.String(),
		u.ID,
	}, "-")
}

func (u *UsersRow) UnmarshalUser(wu *workspace.User) (err error) {
	addresses, err := ParseUserAddresses(wu.Addresses)
	if err != nil {
		return
	}
	u.Addresses = addresses
	u.AgreedToTerms = wu.AgreedToTerms
	u.Aliases = wu.Aliases
	u.Archived = wu.Archived
	u.ChangePasswordAtNextLogin = wu.ChangePasswordAtNextLogin
	u.CreationTime = wu.CreationTime
	u.CustomSchemas = ParseCustomSchemas(wu.CustomSchemas)
	u.CustomerID = wu.CustomerId
	u.DeletionTime = wu.DeletionTime
	emails, err := ParseUserEmails(wu.Emails)
	if err != nil {
		return err
	}
	u.Emails = emails
	u.Etag = wu.Etag
	ids, err := ParseUserExternalIDs(wu.ExternalIds)
	if err != nil {
		return err
	}
	u.ExternalIDs = ids
	gender, err := ParseUserGender(wu.Gender)
	if err != nil {
		return err
	}
	u.Gender = gender
	u.ID = wu.Id
	u.IPWhitelisted = wu.IpWhitelisted
	u.IsAdmin = wu.IsAdmin
	u.IsDelegatedAdmin = wu.IsDelegatedAdmin
	u.IsEnrolledIn2Sv = wu.IsEnrolledIn2Sv
	u.IsEnforcedIn2Sv = wu.IsEnforcedIn2Sv
	u.IsMailboxSetup = wu.IsMailboxSetup
	keywords, err := ParseUserKeywords(wu.Keywords)
	if err != nil {
		return err
	}
	u.Keywords = keywords
	u.Kind = wu.Kind
	languages, err := ParseUserLanguages(wu.Languages)
	if err != nil {
		return err
	}
	u.Languages = languages
	u.LastLoginTime = wu.LastLoginTime
	locations, err := ParseUserLocations(wu.Locations)
	if err != nil {
		return err
	}
	u.Locations = locations
	u.Name.ParseUserName(wu.Name)
	u.NonEditableAliases = wu.NonEditableAliases
	notes, err := ParseUserAbouts(wu.Notes)
	if err != nil {
		return err
	}
	u.Notes = notes
	u.OrgUnitPath = wu.OrgUnitPath
	orgs, err := ParseUserOrganizations(wu.Organizations)
	if err != nil {
		return err
	}
	u.Organizations = orgs
	phones, err := ParseUserPhones(wu.Phones)
	if err != nil {
		return err
	}
	u.Phones = phones
	u.PrimaryEmail = wu.PrimaryEmail
	u.RecoveryEmail = wu.RecoveryEmail
	u.RecoveryPhone = wu.RecoveryPhone
	relations, err := ParseUserRelations(wu.Relations)
	if err != nil {
		return err
	}
	u.Relations = relations
	u.Suspended = wu.Suspended
	u.SuspensionReason = wu.SuspensionReason
	u.ThumbnailPhotoEtag = wu.ThumbnailPhotoEtag
	u.ThumbnailPhotoURL = wu.ThumbnailPhotoUrl
	websites, err := ParseUserWebsites(wu.Websites)
	if err != nil {
		return err
	}
	u.Websites = websites
	return nil
}

func ParseUserAddresses(data interface{}) ([]UserAddress, error) {
	if data == nil {
		return nil, nil
	}
	interfaces := data.([]interface{})
	list := make([]UserAddress, 0)
	for _, face := range interfaces {
		obj := face.(map[string]interface{})
		enc, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		var address UserAddress
		err = json.Unmarshal(enc, &address)
		if err != nil {
			return nil, err
		}
		list = append(list, address)
	}
	return list, nil
}

func ParseUserExternalIDs(data interface{}) ([]UserExternalID, error) {
	if data == nil {
		return nil, nil
	}
	interfaces := data.([]interface{})
	list := make([]UserExternalID, 0, len(interfaces))
	for _, face := range interfaces {
		obj := face.(map[string]interface{})
		var id UserExternalID
		enc, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(enc, &id)
		if err != nil {
			return nil, err
		}
		list = append(list, id)
	}
	return list, nil
}

func ParseUserGender(data interface{}) (UserGender, error) {
	if data == nil {
		return UserGender{}, nil
	}
	obj := data.(map[string]interface{})
	var gender UserGender
	enc, err := json.Marshal(obj)
	if err != nil {
		return UserGender{}, err
	}
	err = json.Unmarshal(enc, &gender)
	if err != nil {
		return UserGender{}, err
	}
	return gender, nil
}

func ParseUserKeywords(data interface{}) ([]UserKeyword, error) {
	if data == nil {
		return nil, nil
	}
	interfaces := data.([]interface{})
	list := make([]UserKeyword, 0, len(interfaces))
	for _, face := range interfaces {
		obj := face.(map[string]interface{})
		enc, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		var keyword UserKeyword
		err = json.Unmarshal(enc, &keyword)
		if err != nil {
			return nil, err
		}
		list = append(list, keyword)
	}
	return list, nil
}

func ParseUserLanguages(data interface{}) ([]UserLanguage, error) {
	if data == nil {
		return nil, nil
	}
	interfaces := data.([]interface{})
	list := make([]UserLanguage, 0, len(interfaces))
	for _, face := range interfaces {
		obj := face.(map[string]interface{})
		enc, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		var language UserLanguage
		err = json.Unmarshal(enc, &language)
		if err != nil {
			return nil, err
		}
		list = append(list, language)
	}
	return list, nil
}

func ParseUserLocations(data interface{}) ([]UserLocation, error) {
	list := make([]UserLocation, 0)
	if data == nil {
		return list, nil
	}
	interfaces := data.([]interface{})
	for _, face := range interfaces {
		obj := face.(map[string]interface{})
		enc, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		var location UserLocation
		err = json.Unmarshal(enc, &location)
		if err != nil {
			return nil, err
		}
		list = append(list, location)
	}
	return list, nil
}

func (u *UserName) ParseUserName(wu *workspace.UserName) {
	u.FamilyName = wu.FamilyName
	u.FullName = wu.FullName
	u.GivenName = wu.GivenName
}

func ParseUserAbouts(data interface{}) ([]UserAbout, error) {
	if data == nil {
		return nil, nil
	}
	interfaces := data.([]interface{})
	list := make([]UserAbout, 0, len(interfaces))
	for _, face := range interfaces {
		obj := face.(map[string]interface{})
		enc, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		var note UserAbout
		if err = json.Unmarshal(enc, &note); err != nil {
			return nil, err
		}
		list = append(list, note)
	}
	return list, nil
}

func ParseUserOrganizations(data interface{}) ([]UserOrganization, error) {
	if data == nil {
		return nil, nil
	}
	interfaces := data.([]interface{})
	list := make([]UserOrganization, 0)
	for _, face := range interfaces {
		obj := face.(map[string]interface{})
		enc, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		var org UserOrganization
		if err = json.Unmarshal(enc, &org); err != nil {
			return nil, err
		}
		list = append(list, org)
	}
	return list, nil
}

func ParseUserPhones(data interface{}) ([]UserPhone, error) {
	if data == nil {
		return nil, nil
	}
	interfaces := data.([]interface{})
	list := make([]UserPhone, 0)
	for _, face := range interfaces {
		obj := face.(map[string]interface{})
		enc, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		var phone UserPhone
		err = json.Unmarshal(enc, &phone)
		if err != nil {
			return nil, err
		}
		list = append(list, phone)
	}
	return list, nil
}

func ParseUserEmails(data interface{}) ([]UserEmail, error) {
	if data == nil {
		return nil, nil
	}
	interfaces := data.([]interface{})
	list := make([]UserEmail, 0, len(interfaces))
	for _, face := range interfaces {
		obj := face.(map[string]interface{})
		enc, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		var email UserEmail
		err = json.Unmarshal(enc, &email)
		if err != nil {
			return nil, err
		}
		list = append(list, email)
	}
	return list, nil
}

func ParseUserRelations(data interface{}) ([]UserRelation, error) {
	if data == nil {
		return nil, nil
	}
	interfaces := data.([]interface{})
	list := make([]UserRelation, 0, len(interfaces))
	for _, face := range interfaces {
		obj := face.(map[string]interface{})
		var relation UserRelation
		enc, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(enc, &relation)
		if err != nil {
			return nil, err
		}
		list = append(list, relation)
	}
	return list, nil
}

func ParseUserWebsites(data interface{}) ([]UserWebsite, error) {
	if data == nil {
		return nil, nil
	}
	interfaces := data.([]interface{})
	list := make([]UserWebsite, 0, len(interfaces))
	for _, face := range interfaces {
		obj := face.(map[string]interface{})
		var website UserWebsite
		enc, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(enc, &website)
		if err != nil {
			return nil, err
		}
		list = append(list, website)
	}
	return list, nil
}

func ParseCustomSchemas(customSchemas map[string]googleapi.RawMessage) []CustomSchema {
	schemas := make([]CustomSchema, 0, len(customSchemas))
	for schema, raw := range customSchemas {
		schemas = append(schemas, CustomSchema{
			Key:   schema,
			Value: string(raw),
		})
	}
	return schemas
}
