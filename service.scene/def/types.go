// Code generated by protoc-gen-jrpc. DO NOT EDIT.

package scenedef

import (
	time "time"

	errors "github.com/jakewright/home-automation/libraries/go/errors"
)

// Scene is defined in the .def file
type Scene struct {
	Id        uint32    `json:"id"`
	Name      string    `json:"name"`
	OwnerId   uint32    `json:"owner_id"`
	Actions   []Action  `json:"actions"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate returns an error if any of the fields have bad values
func (m *Scene) Validate() error {

	for _, r := range m.Actions {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Action is defined in the .def file
type Action struct {
	Stage          int32     `json:"stage"`
	Sequence       int32     `json:"sequence"`
	Func           string    `json:"func"`
	ControllerName string    `json:"controller_name"`
	DeviceId       string    `json:"device_id"`
	Command        string    `json:"command"`
	Property       string    `json:"property"`
	PropertyValue  string    `json:"property_value"`
	PropertyType   string    `json:"property_type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Validate returns an error if any of the fields have bad values
func (m *Action) Validate() error {

	return nil
}

// CreateSceneRequest_Action is defined in the .def file
type CreateSceneRequest_Action struct {
	Stage          int32  `json:"stage"`
	Sequence       int32  `json:"sequence"`
	Func           string `json:"func"`
	ControllerName string `json:"controller_name"`
	DeviceId       string `json:"device_id"`
	Command        string `json:"command"`
	Property       string `json:"property"`
	PropertyValue  string `json:"property_value"`
	PropertyType   string `json:"property_type"`
}

// Validate returns an error if any of the fields have bad values
func (m *CreateSceneRequest_Action) Validate() error {

	return nil
}

// CreateSceneRequest is defined in the .def file
type CreateSceneRequest struct {
	Name    string                      `json:"name"`
	OwnerId uint32                      `json:"owner_id"`
	Actions []CreateSceneRequest_Action `json:"actions"`
}

// Validate returns an error if any of the fields have bad values
func (m *CreateSceneRequest) Validate() error {
	if m.Name == "" {
		return errors.BadRequest("field name is required")
	}

	if m.OwnerId == 0 {
		return errors.BadRequest("field owner_id is required")
	}

	for _, r := range m.Actions {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	if len(m.Actions) == 0 {
		return errors.BadRequest("field actions is required")
	}

	return nil
}

// CreateSceneResponse is defined in the .def file
type CreateSceneResponse struct {
	Scene Scene `json:"scene"`
}

// Validate returns an error if any of the fields have bad values
func (m *CreateSceneResponse) Validate() error {
	if err := m.Scene.Validate(); err != nil {
		return err
	}

	return nil
}

// ReadSceneRequest is defined in the .def file
type ReadSceneRequest struct {
	SceneId uint32 `json:"scene_id"`
}

// Validate returns an error if any of the fields have bad values
func (m *ReadSceneRequest) Validate() error {

	return nil
}

// ReadSceneResponse is defined in the .def file
type ReadSceneResponse struct {
	Scene Scene `json:"scene"`
}

// Validate returns an error if any of the fields have bad values
func (m *ReadSceneResponse) Validate() error {
	if err := m.Scene.Validate(); err != nil {
		return err
	}

	return nil
}

// ListScenesRequest is defined in the .def file
type ListScenesRequest struct {
	OwnerId uint32 `json:"owner_id"`
}

// Validate returns an error if any of the fields have bad values
func (m *ListScenesRequest) Validate() error {

	return nil
}

// ListScenesResponse is defined in the .def file
type ListScenesResponse struct {
	Scenes []Scene `json:"scenes"`
}

// Validate returns an error if any of the fields have bad values
func (m *ListScenesResponse) Validate() error {
	for _, r := range m.Scenes {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// DeleteSceneRequest is defined in the .def file
type DeleteSceneRequest struct {
	SceneId int32 `json:"scene_id"`
}

// Validate returns an error if any of the fields have bad values
func (m *DeleteSceneRequest) Validate() error {

	return nil
}

// DeleteSceneResponse is defined in the .def file
type DeleteSceneResponse struct {
}

// Validate returns an error if any of the fields have bad values
func (m *DeleteSceneResponse) Validate() error {
	return nil
}

// SetSceneRequest is defined in the .def file
type SetSceneRequest struct {
	SceneId uint32 `json:"scene_id"`
}

// Validate returns an error if any of the fields have bad values
func (m *SetSceneRequest) Validate() error {

	return nil
}

// SetSceneResponse is defined in the .def file
type SetSceneResponse struct {
}

// Validate returns an error if any of the fields have bad values
func (m *SetSceneResponse) Validate() error {
	return nil
}

// SetSceneEvent is defined in the .def file
type SetSceneEvent struct {
	SceneId uint32 `json:"scene_id"`
}

// Validate returns an error if any of the fields have bad values
func (m *SetSceneEvent) Validate() error {

	return nil
}
