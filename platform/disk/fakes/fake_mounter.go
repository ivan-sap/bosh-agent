package fakes

type FakeMounter struct {
	MountCalled         bool
	MountPartitionPaths []string
	MountMountPoints    []string
	MountMountOptions   [][]string
	MountErr            error

	RemountInPlaceCalled       bool
	RemountInPlaceMountPoint   string
	RemountInPlaceMountOptions []string
	RemountInPlaceErr          error

	RemountAsReadonlyCalled bool
	RemountAsReadonlyPath   string
	RemountAsReadonlyErr    error

	RemountFromMountPoint string
	RemountToMountPoint   string
	RemountMountOptions   []string
	RemountErr            error

	SwapOnPartitionPaths []string
	SwapOnErr            error

	UnmountPartitionPathOrMountPoint string
	UnmountDidUnmount                bool
	UnmountErr                       error

	IsMountPointPath          string
	IsMountPointPartitionPath string
	IsMountPointResult        bool
	IsMountPointErr           error

	IsMountedResult bool
	IsMountedErr    error
	IsMountedStub   func(string) (bool, error)
	isMountedArgs   []string
}

func (m *FakeMounter) Mount(partitionPath, mountPoint string, mountOptions ...string) error {
	m.MountCalled = true
	m.MountPartitionPaths = append(m.MountPartitionPaths, partitionPath)
	m.MountMountPoints = append(m.MountMountPoints, mountPoint)
	m.MountMountOptions = append(m.MountMountOptions, mountOptions)
	return m.MountErr
}

func (m *FakeMounter) RemountAsReadonly(mountPoint string) (err error) {
	m.RemountAsReadonlyCalled = true
	m.RemountAsReadonlyPath = mountPoint
	return m.RemountAsReadonlyErr
}

func (m *FakeMounter) Remount(fromMountPoint, toMountPoint string, mountOptions ...string) (err error) {
	m.RemountFromMountPoint = fromMountPoint
	m.RemountToMountPoint = toMountPoint
	m.RemountMountOptions = mountOptions
	return m.RemountErr
}

func (m *FakeMounter) SwapOn(partitionPath string) (err error) {
	m.SwapOnPartitionPaths = append(m.SwapOnPartitionPaths, partitionPath)
	return m.SwapOnErr
}

func (m *FakeMounter) Unmount(partitionPathOrMountPoint string) (didUnmount bool, err error) {
	m.UnmountPartitionPathOrMountPoint = partitionPathOrMountPoint
	return m.UnmountDidUnmount, m.UnmountErr
}

func (m *FakeMounter) IsMountPoint(path string) (partitionPath string, result bool, err error) {
	m.IsMountPointPath = path
	return m.IsMountPointPartitionPath, m.IsMountPointResult, m.IsMountPointErr
}

func (m *FakeMounter) IsMounted(devicePathOrMountPoint string) (bool, error) {
	m.isMountedArgs = append(m.isMountedArgs, devicePathOrMountPoint)
	if m.IsMountedStub != nil {
		return m.IsMountedStub(devicePathOrMountPoint)
	}
	return m.IsMountedResult, m.IsMountedErr
}

func (m *FakeMounter) IsMountedArgsForCall(callNumber int) string {
	return m.isMountedArgs[callNumber]
}

func (m *FakeMounter) RemountInPlace(mountPoint string, mountOptions ...string) error {
	m.RemountInPlaceCalled = true
	m.RemountInPlaceMountPoint = mountPoint
	m.RemountInPlaceMountOptions = mountOptions
	return m.RemountInPlaceErr
}
