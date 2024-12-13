package peerflow

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/api/workflowservicemock/v1"
	"go.temporal.io/sdk/worker"
	"testing"
)

type replayTestSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
	service  *workflowservicemock.MockWorkflowServiceClient
}

func TestReplayTestSuite(t *testing.T) {
	s := new(replayTestSuite)
	suite.Run(t, s)
}

func (s *replayTestSuite) SetupTest() {
	s.mockCtrl = gomock.NewController(s.T())
	s.service = workflowservicemock.NewMockWorkflowServiceClient(s.mockCtrl)
}

func (s *replayTestSuite) TearDownTest() {
	s.mockCtrl.Finish() // assert mock’s expectations
}

// You can use a replay test to examine a previous execution
//
// "*_events.json" can be downloaded from Temporal CLI, or from Temporal Web UI.
//
//	tctl wf show -w hello_world_workflowID --output_filename ./events.json
//
// Make sure to download the encoded format of the history if you're using the UI.
func (s *replayTestSuite) TestReplayWorkflowHistoryFromFile() {
	replayer := worker.NewWorkflowReplayer()

	replayer.RegisterWorkflow(CDCFlowWorkflow)

	err := replayer.ReplayWorkflowHistoryFromJSONFile(nil, "events.json")
	require.NoError(s.T(), err)
}
