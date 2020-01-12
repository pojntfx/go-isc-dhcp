package utils

import (
	"os/exec"
	"syscall"
)

// ProcessWorker is a worker that manages a process.
type ProcessWorker struct {
	Instance             *exec.Cmd
	ScheduledForDeletion bool
}

// Wait waits for the worker to stop.
func (w *ProcessWorker) Wait() error {
	_, err := w.Instance.Process.Wait()

	return err
}

// DisableAutoRestart disables the auto restart if the process exits.
func (w *ProcessWorker) DisableAutoRestart() error {
	w.ScheduledForDeletion = true

	return nil
}

// Stop stops the process.
func (w *ProcessWorker) Stop() error {
	if err := w.DisableAutoRestart(); err != nil {
		return err
	}

	processGroupID, err := syscall.Getpgid(w.Instance.Process.Pid)
	if err != nil {
		return err
	}

	if err := syscall.Kill(processGroupID, syscall.SIGKILL); err != nil {
		return err
	}

	return nil
}

// IsScheduledForDeletion returns true if the process is scheduled for deletion.
func (w *ProcessWorker) IsScheduledForDeletion() bool {
	return w.ScheduledForDeletion
}

// IsRunning returns true if the process is still running.
func (w *ProcessWorker) IsRunning() bool {
	return w.Instance.Process != nil
}
