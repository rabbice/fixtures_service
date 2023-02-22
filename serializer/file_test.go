package serializer

import (
	"testing"

	pb "github.com/rabbice/fixturesbook/pb"
	"github.com/rabbice/fixturesbook/sample"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/fixtures.bin"
	jsonFile := "../tmp/fixtures.json"

	fixture1 := sample.NewFixture()
	err := WriteProtobufToBinaryFile(fixture1, binaryFile)
	require.NoError(t, err)

	fixture2 := &pb.Fixture{}
	err = ReadProtobufFromBinaryFile(binaryFile, fixture2)
	require.NoError(t, err)
	require.True(t, proto.Equal(fixture1, fixture2))

	err = WriteProtobufToJSONFile(fixture1, jsonFile)
	require.NoError(t, err)
}
