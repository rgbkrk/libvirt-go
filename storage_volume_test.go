package libvirt

import (
	"time"
)

func testStorageVolXML(volName, poolName string) string {
	defName := volName
	if defName == "" {
		defName = time.Now().String()
	}

	defName += ".img"
	return `<volume>
        <name>` + defName + `</name>
        <allocation>0</allocation>
        <capacity unit="M">10</capacity>
        <target>
          <path>` + poolName + "/" + defName + `</path>
          <permissions>
            <owner>107</owner>
            <group>107</group>
            <mode>0744</mode>
            <label>testLabel0</label>
          </permissions>
        </target>
      </volume>`
}
