package osexec

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
)

type OsExecHandler struct {
}

type Spec struct {
	Command []string `json:"command"`

	Env    map[string]string `json:"data"`
	Output *Output           `json:"output"`
}

type Output struct {
	Mask      bool     `json:"mask"`
	Variables []string `json:"variables"`
}

func (h *OsExecHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	in, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO: proper error message
		return
	}
	spec := &Spec{}
	if err := json.Unmarshal(in, spec); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// TODO: proper error message
		return
	}

	// TODO: resolve expression in the ENV. The expression should be only about secrets

	// execute the command
	if err := runCmd(r.Context(), spec.Command, spec.Env); err != nil {
		// TODO: we should design Runner level general error types. Based on error type, put different http response code and error message
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// collect outputs

	w.WriteHeader(http.StatusOK)
	return
}

func runCmd(ctx context.Context, command []string, envs map[string]string) error {
	if len(command) == 0 {
		return errors.New("command cannot be empty")
	}
	startTime := time.Now()
	cmdArgs := make([]string, 0)
	cmdArgs = append(cmdArgs, command[1:]...)

	cmd := exec.Command(command[0], cmdArgs...) //nolint:gosec
	// TODO: set workdir to a tmp folder. eg. cmd.Dir = ${CURRENT}/taskId
	cmd.Env = toEnv(envs)
	// TODO: redirect output to proper channel, eg. logstream logger
	// cmd.Stderr = output
	// cmd.Stdout = output
	err := cmd.Start()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error":           err,
			"commands":        cmd,
			"elapsed_time_ms": time.Since(startTime),
		}).Error("starting command failed")
		return err
	}

	err = cmd.Wait()

	if ctxErr := ctx.Err(); ctxErr == context.DeadlineExceeded {
		logrus.WithFields(logrus.Fields{
			"commands":        cmd,
			"elapsed_time_ms": time.Since(startTime),
		}).Error("timeout while executing the step")

		return ctxErr
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error":           err,
			"commands":        cmd,
			"elapsed_time_ms": time.Since(startTime),
		}).Error("error encountered while executing the step")
		return err
	}
	return nil
}

// helper function that converts a key value map of
// environment variables to a string slice in key=value
// format.
func toEnv(env map[string]string) []string {
	var envs []string
	for k, v := range env {
		if v != "" {
			envs = append(envs, k+"="+v)
		}
	}
	return envs
}
