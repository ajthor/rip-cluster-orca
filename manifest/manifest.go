package manifest

import (
  "reflect"
  // "bytes"
  // "strings"

  // "html/template"
  // "net/url"
  // "path/filepath"

  // "github.com/Masterminds/semver"
  // log "github.com/gorobot/robologger"
)

// Generated using: https://mholt.github.io/json-to-go/
type Manifest struct {
  Name string `json:"name"`
  Repo string `json:"repo"`
  Registry string `json:"registry"`
  Images []Image `json:"images"`
}

// NewManifest creates a new, empty manifest.
func NewManifest() *Manifest {
  return &Manifest{}
}

// Get returns the image in the manifest corresponding with the supplied name.
// Returns a nil pointer if no manifest matching the provided name can be
// found.
func (m Manifest) Get(name string) *Image {
  for i, img := range m.Images {
    if name == img.Name {
      return &m.Images[i]
    }
  }

  return nil
}

// Filter filters the images in the manifest and returns the matches as a
// slice. If no matches are found, the function returns an empty slice.
func (m Manifest) Filter(f func(Image) bool) []*Image {
  matches := make([]*Image, 0)
  for i, img := range m.Images {
    if f(img) {
      matches = append(matches, &m.Images[i])
    }
  }

  return matches
}

// Match returns the first matching manifest based on the constraints specified
// by the map[string]interface{} argument. It returns the first image found.
func (m Manifest) Match(args map[string]interface{}) *Image {
  // Iterate through all images.
  for i, img := range m.Images {
    val := reflect.ValueOf(img)

    // Go through the constraints, one by one.
    for k, v := range args {
      field := val.FieldByName(k)
      if !field.IsValid() {
        break
      }

      if field.Kind() == reflect.Slice {
        for j := 0; j < field.Len(); j++ {
          if reflect.DeepEqual(v, field.Index(j).Interface()) {
            return &m.Images[i]
          }
        }
        break
      }

      if reflect.DeepEqual(v, field.Interface()) {
        return &m.Images[i]
      }
    }
  }

  return nil
}

// MatchAll returns the first matching manifest based on the constraints
// specified by the map[string]interface{} argument. It returns all images
// found.
func (m Manifest) MatchAll(args map[string]interface{}) []*Image {
  matches := make([]*Image, 0)
  // Iterate through all images.
  for i, img := range m.Images {
    val := reflect.ValueOf(img)

    // Go through the constraints, one by one.
    for k, v := range args {
      field := val.FieldByName(k)
      if !field.IsValid() {
        break
      }

      if field.Kind() == reflect.Slice {
        for j := 0; j < field.Len(); j++ {
          if reflect.DeepEqual(v, field.Index(j).Interface()) {
            matches = append(matches, &m.Images[i])
          }
        }
        break
      }

      if reflect.DeepEqual(v, field.Interface()) {
        matches = append(matches, &m.Images[i])
      }
    }
  }

  return matches
}

// // GetTags returns all of the tags associated with the maniefest. It generates
// // the tags using the registry, the repo, the image name, and the version.
// func (m *Manifest) GetTags() []string {
//   tags := []string{}
//   names := m.GetImageNames()
//
//   // Generate the build tags.
//   for _, v := range m.Versions {
//     // Append the version to the image name.
//     for _, t := range names {
//       s := []string{t, v}
//       tags = append(tags, strings.Join(s, ":"))
//     }
//   }
//
//   return tags
// }

// It generates the names using the registry, repo, and name found in the
// manifest.
// func (m *Manifest) GetImageNames() []string {
//   names = make([]string, 0)
//
//   if m.Repo != "" {
//     s := []string{m.Repo, m.Name}
//     names = append(names, strings.Join(s, "/"))
//   } else {
//     names = append(names, m.Name)
//   }
//
//   if m.Registry != "" {
//     for _, t := range names {
//       s := []string{m.Registry, t}
//       names = append(names, strings.Join(s, "/"))
//     }
//   }
//
//   return names
// }
