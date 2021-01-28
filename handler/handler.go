package handler

import (
	"errors"
	"strconv"

	gq "github.com/eiri/goldberg/queue"
)

type Request struct {
	Name string
	Item string
}

type Response struct {
	Message string
}

type Handler struct {
	dispatch map[string]gq.Queue
}

func New() *Handler {
	return &Handler{dispatch: make(map[string]gq.Queue)}
}

func (h *Handler) Create(req Request, resp *Response) error {
	if req.Name == "" {
		return errors.New("Missing queue name")
	}

	h.dispatch[req.Name] = gq.NewFIFO()

	resp.Message = "ok"
	return nil
}

func (h *Handler) Delete(req Request, resp *Response) error {
	return nil
}

func (h *Handler) PushBack(req Request, resp *Response) error {
	if req.Name == "" {
		return errors.New("Missing queue name")
	}
	if req.Item == "" {
		return errors.New("Missing item")
	}

	h.dispatch[req.Name].PushBack(req.Item)

	resp.Message = "ok"
	return nil
}

func (h *Handler) PushFront(req Request, resp *Response) error {
	return nil
}

func (h *Handler) PopBack(req Request, resp *Response) error {
	return nil
}

func (h *Handler) PopFront(req Request, resp *Response) error {
	if req.Name == "" {
		return errors.New("Missing queue name")
	}

	item := h.dispatch[req.Name].PopFront()

	e, ok := item.(string)
	if !ok {
		return errors.New("Invalid element type")
	}

	resp.Message = e
	return nil
}

func (h *Handler) Back(req Request, resp *Response) error {
	return nil
}

func (h *Handler) Front(req Request, resp *Response) error {
	return nil
}

func (h *Handler) Len(req Request, resp *Response) error {
	if req.Name == "" {
		return errors.New("Missing queue name")
	}

	len := h.dispatch[req.Name].Len()

	resp.Message = strconv.Itoa(len)
	return nil
}
