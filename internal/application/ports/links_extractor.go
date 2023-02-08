package ports

import (
	"golang-web-crawler/internal/application/models"
	"io"
)

type LinksExtractor interface {
	All(htmlBody io.Reader) ([]models.Link, error)
}
