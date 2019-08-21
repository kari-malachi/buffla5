package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/kari-malachi/buffla5/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Link)
// DB Table: Plural (links)
// Resource: Plural (Links)
// Path: Plural (/links)
// View Template Folder: Plural (/templates/links/)

// LinksResource is the resource for the Link model
type LinksResource struct {
	buffalo.Resource
}

// List gets all Links. This function is mapped to the path
// GET /links
func (v LinksResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	links := &models.Links{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Links from the DB
	if err := q.All(links); err != nil {
		return err
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, links))
}

// Show gets the data for one Link. This function is mapped to
// the path GET /links/{link_id}
func (v LinksResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Link
	link := &models.Link{}

	// To find the Link the parameter link_id is used.
	if err := tx.Find(link, c.Param("link_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, link))
}

// New renders the form for creating a new Link.
// This function is mapped to the path GET /links/new
func (v LinksResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Link{}))
}

// Create adds a Link to the DB. This function is mapped to the
// path POST /links
func (v LinksResource) Create(c buffalo.Context) error {
	// Allocate an empty Link
	link := &models.Link{}

	// Bind link to the html form elements
	if err := c.Bind(link); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(link)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, link))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "link.created.success"))
	// and redirect to the links index page
	return c.Render(201, r.Auto(c, link))
}

// Edit renders a edit form for a Link. This function is
// mapped to the path GET /links/{link_id}/edit
func (v LinksResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Link
	link := &models.Link{}

	if err := tx.Find(link, c.Param("link_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, link))
}

// Update changes a Link in the DB. This function is mapped to
// the path PUT /links/{link_id}
func (v LinksResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Link
	link := &models.Link{}

	if err := tx.Find(link, c.Param("link_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Link to the html form elements
	if err := c.Bind(link); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(link)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, link))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "link.updated.success"))
	// and redirect to the links index page
	return c.Render(200, r.Auto(c, link))
}

// Destroy deletes a Link from the DB. This function is mapped
// to the path DELETE /links/{link_id}
func (v LinksResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Link
	link := &models.Link{}

	// To find the Link the parameter link_id is used.
	if err := tx.Find(link, c.Param("link_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(link); err != nil {
		return err
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", T.Translate(c, "link.destroyed.success"))
	// Redirect to the links index page
	return c.Render(200, r.Auto(c, link))
}
