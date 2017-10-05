package obj

import (
  "time"
  "github.com/Alluxio/alluxio-go"
)

func NewAlluxioClient(host string, port int, timeout time.Duration) (Client, error) {
  return &alluxioClient{fs: alluxio.NewClient(host, port, timeout)}
}

type alluxioClient struct {
  fs alluxio.Client
}

func (c *alluxioClient) Writer(path string) (io.WriteCloser, error) {
}

func (c *alluxioClient) Reader(path string, offset uint64, size uint64) (io.ReadCloser, error) {
}

func (c *alluxioClient) Delete(path string) error {
}

func (c *alluxioClient) Walk(dir string, walkFn func(name string) error) error {
}

func (c *alluxioClient) Exists(path string) bool {
}

func (c *alluxioClient) IsRetryable(err error) bool {
	return false
}

func (c *alluxioClient) IsNotExist(err error) bool {
	return false
}

func (c *alluxioClient) IsIgnorable(err error) bool {
	return false
}
