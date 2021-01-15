# test-gofs
Test project to test gofsutil Bindmount method.

Create a pvc and a pod - make sure the pod is in running state
Extract staging path from node logs or by using `mount | grep csi`

Captrue the output of `mount | grep csi`

Download the project and build using `go build`
A binary should be generated test-gofs
Create a new test target directory like /root/user/csi-target
Now run the binary using arguments -staging-path and -target-path

Example: ./test-gofs -staging-path <stagingpath> -target-path /root/user/csi-target

Successful execution will have the below output
Performin Bind of staging to target pathINFO[0000] mount command                                 args="-o bind <stagingpath> /root/user/csi-target" cmd=mount
INFO[0000] mount command                                 args="-o remount <stagingpath> /root/user/csi-target" cmd=mount


Captrue the output of `mount | grep csi` again

The latest output of mount should contain /root/user/csi-target
