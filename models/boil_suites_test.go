// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Auths", testAuths)
	t.Run("Changelogs", testChangelogs)
	t.Run("Confidences", testConfidences)
	t.Run("Containers", testContainers)
	t.Run("Cpes", testCpes)
	t.Run("CveDetails", testCveDetails)
	t.Run("DistroAdvisories", testDistroAdvisories)
	t.Run("Groups", testGroups)
	t.Run("JVNS", testJVNS)
	t.Run("NVDS", testNVDS)
	t.Run("Organizations", testOrganizations)
	t.Run("PackageInfos", testPackageInfos)
	t.Run("Platforms", testPlatforms)
	t.Run("References", testReferences)
	t.Run("ScanResults", testScanResults)
	t.Run("Tasks", testTasks)
	t.Run("Users", testUsers)
	t.Run("VulnInfos", testVulnInfos)
}

func TestDelete(t *testing.T) {
	t.Run("Auths", testAuthsDelete)
	t.Run("Changelogs", testChangelogsDelete)
	t.Run("Confidences", testConfidencesDelete)
	t.Run("Containers", testContainersDelete)
	t.Run("Cpes", testCpesDelete)
	t.Run("CveDetails", testCveDetailsDelete)
	t.Run("DistroAdvisories", testDistroAdvisoriesDelete)
	t.Run("Groups", testGroupsDelete)
	t.Run("JVNS", testJVNSDelete)
	t.Run("NVDS", testNVDSDelete)
	t.Run("Organizations", testOrganizationsDelete)
	t.Run("PackageInfos", testPackageInfosDelete)
	t.Run("Platforms", testPlatformsDelete)
	t.Run("References", testReferencesDelete)
	t.Run("ScanResults", testScanResultsDelete)
	t.Run("Tasks", testTasksDelete)
	t.Run("Users", testUsersDelete)
	t.Run("VulnInfos", testVulnInfosDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Auths", testAuthsQueryDeleteAll)
	t.Run("Changelogs", testChangelogsQueryDeleteAll)
	t.Run("Confidences", testConfidencesQueryDeleteAll)
	t.Run("Containers", testContainersQueryDeleteAll)
	t.Run("Cpes", testCpesQueryDeleteAll)
	t.Run("CveDetails", testCveDetailsQueryDeleteAll)
	t.Run("DistroAdvisories", testDistroAdvisoriesQueryDeleteAll)
	t.Run("Groups", testGroupsQueryDeleteAll)
	t.Run("JVNS", testJVNSQueryDeleteAll)
	t.Run("NVDS", testNVDSQueryDeleteAll)
	t.Run("Organizations", testOrganizationsQueryDeleteAll)
	t.Run("PackageInfos", testPackageInfosQueryDeleteAll)
	t.Run("Platforms", testPlatformsQueryDeleteAll)
	t.Run("References", testReferencesQueryDeleteAll)
	t.Run("ScanResults", testScanResultsQueryDeleteAll)
	t.Run("Tasks", testTasksQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
	t.Run("VulnInfos", testVulnInfosQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Auths", testAuthsSliceDeleteAll)
	t.Run("Changelogs", testChangelogsSliceDeleteAll)
	t.Run("Confidences", testConfidencesSliceDeleteAll)
	t.Run("Containers", testContainersSliceDeleteAll)
	t.Run("Cpes", testCpesSliceDeleteAll)
	t.Run("CveDetails", testCveDetailsSliceDeleteAll)
	t.Run("DistroAdvisories", testDistroAdvisoriesSliceDeleteAll)
	t.Run("Groups", testGroupsSliceDeleteAll)
	t.Run("JVNS", testJVNSSliceDeleteAll)
	t.Run("NVDS", testNVDSSliceDeleteAll)
	t.Run("Organizations", testOrganizationsSliceDeleteAll)
	t.Run("PackageInfos", testPackageInfosSliceDeleteAll)
	t.Run("Platforms", testPlatformsSliceDeleteAll)
	t.Run("References", testReferencesSliceDeleteAll)
	t.Run("ScanResults", testScanResultsSliceDeleteAll)
	t.Run("Tasks", testTasksSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
	t.Run("VulnInfos", testVulnInfosSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Auths", testAuthsExists)
	t.Run("Changelogs", testChangelogsExists)
	t.Run("Confidences", testConfidencesExists)
	t.Run("Containers", testContainersExists)
	t.Run("Cpes", testCpesExists)
	t.Run("CveDetails", testCveDetailsExists)
	t.Run("DistroAdvisories", testDistroAdvisoriesExists)
	t.Run("Groups", testGroupsExists)
	t.Run("JVNS", testJVNSExists)
	t.Run("NVDS", testNVDSExists)
	t.Run("Organizations", testOrganizationsExists)
	t.Run("PackageInfos", testPackageInfosExists)
	t.Run("Platforms", testPlatformsExists)
	t.Run("References", testReferencesExists)
	t.Run("ScanResults", testScanResultsExists)
	t.Run("Tasks", testTasksExists)
	t.Run("Users", testUsersExists)
	t.Run("VulnInfos", testVulnInfosExists)
}

func TestFind(t *testing.T) {
	t.Run("Auths", testAuthsFind)
	t.Run("Changelogs", testChangelogsFind)
	t.Run("Confidences", testConfidencesFind)
	t.Run("Containers", testContainersFind)
	t.Run("Cpes", testCpesFind)
	t.Run("CveDetails", testCveDetailsFind)
	t.Run("DistroAdvisories", testDistroAdvisoriesFind)
	t.Run("Groups", testGroupsFind)
	t.Run("JVNS", testJVNSFind)
	t.Run("NVDS", testNVDSFind)
	t.Run("Organizations", testOrganizationsFind)
	t.Run("PackageInfos", testPackageInfosFind)
	t.Run("Platforms", testPlatformsFind)
	t.Run("References", testReferencesFind)
	t.Run("ScanResults", testScanResultsFind)
	t.Run("Tasks", testTasksFind)
	t.Run("Users", testUsersFind)
	t.Run("VulnInfos", testVulnInfosFind)
}

func TestBind(t *testing.T) {
	t.Run("Auths", testAuthsBind)
	t.Run("Changelogs", testChangelogsBind)
	t.Run("Confidences", testConfidencesBind)
	t.Run("Containers", testContainersBind)
	t.Run("Cpes", testCpesBind)
	t.Run("CveDetails", testCveDetailsBind)
	t.Run("DistroAdvisories", testDistroAdvisoriesBind)
	t.Run("Groups", testGroupsBind)
	t.Run("JVNS", testJVNSBind)
	t.Run("NVDS", testNVDSBind)
	t.Run("Organizations", testOrganizationsBind)
	t.Run("PackageInfos", testPackageInfosBind)
	t.Run("Platforms", testPlatformsBind)
	t.Run("References", testReferencesBind)
	t.Run("ScanResults", testScanResultsBind)
	t.Run("Tasks", testTasksBind)
	t.Run("Users", testUsersBind)
	t.Run("VulnInfos", testVulnInfosBind)
}

func TestOne(t *testing.T) {
	t.Run("Auths", testAuthsOne)
	t.Run("Changelogs", testChangelogsOne)
	t.Run("Confidences", testConfidencesOne)
	t.Run("Containers", testContainersOne)
	t.Run("Cpes", testCpesOne)
	t.Run("CveDetails", testCveDetailsOne)
	t.Run("DistroAdvisories", testDistroAdvisoriesOne)
	t.Run("Groups", testGroupsOne)
	t.Run("JVNS", testJVNSOne)
	t.Run("NVDS", testNVDSOne)
	t.Run("Organizations", testOrganizationsOne)
	t.Run("PackageInfos", testPackageInfosOne)
	t.Run("Platforms", testPlatformsOne)
	t.Run("References", testReferencesOne)
	t.Run("ScanResults", testScanResultsOne)
	t.Run("Tasks", testTasksOne)
	t.Run("Users", testUsersOne)
	t.Run("VulnInfos", testVulnInfosOne)
}

func TestAll(t *testing.T) {
	t.Run("Auths", testAuthsAll)
	t.Run("Changelogs", testChangelogsAll)
	t.Run("Confidences", testConfidencesAll)
	t.Run("Containers", testContainersAll)
	t.Run("Cpes", testCpesAll)
	t.Run("CveDetails", testCveDetailsAll)
	t.Run("DistroAdvisories", testDistroAdvisoriesAll)
	t.Run("Groups", testGroupsAll)
	t.Run("JVNS", testJVNSAll)
	t.Run("NVDS", testNVDSAll)
	t.Run("Organizations", testOrganizationsAll)
	t.Run("PackageInfos", testPackageInfosAll)
	t.Run("Platforms", testPlatformsAll)
	t.Run("References", testReferencesAll)
	t.Run("ScanResults", testScanResultsAll)
	t.Run("Tasks", testTasksAll)
	t.Run("Users", testUsersAll)
	t.Run("VulnInfos", testVulnInfosAll)
}

func TestCount(t *testing.T) {
	t.Run("Auths", testAuthsCount)
	t.Run("Changelogs", testChangelogsCount)
	t.Run("Confidences", testConfidencesCount)
	t.Run("Containers", testContainersCount)
	t.Run("Cpes", testCpesCount)
	t.Run("CveDetails", testCveDetailsCount)
	t.Run("DistroAdvisories", testDistroAdvisoriesCount)
	t.Run("Groups", testGroupsCount)
	t.Run("JVNS", testJVNSCount)
	t.Run("NVDS", testNVDSCount)
	t.Run("Organizations", testOrganizationsCount)
	t.Run("PackageInfos", testPackageInfosCount)
	t.Run("Platforms", testPlatformsCount)
	t.Run("References", testReferencesCount)
	t.Run("ScanResults", testScanResultsCount)
	t.Run("Tasks", testTasksCount)
	t.Run("Users", testUsersCount)
	t.Run("VulnInfos", testVulnInfosCount)
}

func TestHooks(t *testing.T) {
	t.Run("Auths", testAuthsHooks)
	t.Run("Changelogs", testChangelogsHooks)
	t.Run("Confidences", testConfidencesHooks)
	t.Run("Containers", testContainersHooks)
	t.Run("Cpes", testCpesHooks)
	t.Run("CveDetails", testCveDetailsHooks)
	t.Run("DistroAdvisories", testDistroAdvisoriesHooks)
	t.Run("Groups", testGroupsHooks)
	t.Run("JVNS", testJVNSHooks)
	t.Run("NVDS", testNVDSHooks)
	t.Run("Organizations", testOrganizationsHooks)
	t.Run("PackageInfos", testPackageInfosHooks)
	t.Run("Platforms", testPlatformsHooks)
	t.Run("References", testReferencesHooks)
	t.Run("ScanResults", testScanResultsHooks)
	t.Run("Tasks", testTasksHooks)
	t.Run("Users", testUsersHooks)
	t.Run("VulnInfos", testVulnInfosHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Auths", testAuthsInsert)
	t.Run("Auths", testAuthsInsertWhitelist)
	t.Run("Changelogs", testChangelogsInsert)
	t.Run("Changelogs", testChangelogsInsertWhitelist)
	t.Run("Confidences", testConfidencesInsert)
	t.Run("Confidences", testConfidencesInsertWhitelist)
	t.Run("Containers", testContainersInsert)
	t.Run("Containers", testContainersInsertWhitelist)
	t.Run("Cpes", testCpesInsert)
	t.Run("Cpes", testCpesInsertWhitelist)
	t.Run("CveDetails", testCveDetailsInsert)
	t.Run("CveDetails", testCveDetailsInsertWhitelist)
	t.Run("DistroAdvisories", testDistroAdvisoriesInsert)
	t.Run("DistroAdvisories", testDistroAdvisoriesInsertWhitelist)
	t.Run("Groups", testGroupsInsert)
	t.Run("Groups", testGroupsInsertWhitelist)
	t.Run("JVNS", testJVNSInsert)
	t.Run("JVNS", testJVNSInsertWhitelist)
	t.Run("NVDS", testNVDSInsert)
	t.Run("NVDS", testNVDSInsertWhitelist)
	t.Run("Organizations", testOrganizationsInsert)
	t.Run("Organizations", testOrganizationsInsertWhitelist)
	t.Run("PackageInfos", testPackageInfosInsert)
	t.Run("PackageInfos", testPackageInfosInsertWhitelist)
	t.Run("Platforms", testPlatformsInsert)
	t.Run("Platforms", testPlatformsInsertWhitelist)
	t.Run("References", testReferencesInsert)
	t.Run("References", testReferencesInsertWhitelist)
	t.Run("ScanResults", testScanResultsInsert)
	t.Run("ScanResults", testScanResultsInsertWhitelist)
	t.Run("Tasks", testTasksInsert)
	t.Run("Tasks", testTasksInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
	t.Run("VulnInfos", testVulnInfosInsert)
	t.Run("VulnInfos", testVulnInfosInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Auths", testAuthsReload)
	t.Run("Changelogs", testChangelogsReload)
	t.Run("Confidences", testConfidencesReload)
	t.Run("Containers", testContainersReload)
	t.Run("Cpes", testCpesReload)
	t.Run("CveDetails", testCveDetailsReload)
	t.Run("DistroAdvisories", testDistroAdvisoriesReload)
	t.Run("Groups", testGroupsReload)
	t.Run("JVNS", testJVNSReload)
	t.Run("NVDS", testNVDSReload)
	t.Run("Organizations", testOrganizationsReload)
	t.Run("PackageInfos", testPackageInfosReload)
	t.Run("Platforms", testPlatformsReload)
	t.Run("References", testReferencesReload)
	t.Run("ScanResults", testScanResultsReload)
	t.Run("Tasks", testTasksReload)
	t.Run("Users", testUsersReload)
	t.Run("VulnInfos", testVulnInfosReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Auths", testAuthsReloadAll)
	t.Run("Changelogs", testChangelogsReloadAll)
	t.Run("Confidences", testConfidencesReloadAll)
	t.Run("Containers", testContainersReloadAll)
	t.Run("Cpes", testCpesReloadAll)
	t.Run("CveDetails", testCveDetailsReloadAll)
	t.Run("DistroAdvisories", testDistroAdvisoriesReloadAll)
	t.Run("Groups", testGroupsReloadAll)
	t.Run("JVNS", testJVNSReloadAll)
	t.Run("NVDS", testNVDSReloadAll)
	t.Run("Organizations", testOrganizationsReloadAll)
	t.Run("PackageInfos", testPackageInfosReloadAll)
	t.Run("Platforms", testPlatformsReloadAll)
	t.Run("References", testReferencesReloadAll)
	t.Run("ScanResults", testScanResultsReloadAll)
	t.Run("Tasks", testTasksReloadAll)
	t.Run("Users", testUsersReloadAll)
	t.Run("VulnInfos", testVulnInfosReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Auths", testAuthsSelect)
	t.Run("Changelogs", testChangelogsSelect)
	t.Run("Confidences", testConfidencesSelect)
	t.Run("Containers", testContainersSelect)
	t.Run("Cpes", testCpesSelect)
	t.Run("CveDetails", testCveDetailsSelect)
	t.Run("DistroAdvisories", testDistroAdvisoriesSelect)
	t.Run("Groups", testGroupsSelect)
	t.Run("JVNS", testJVNSSelect)
	t.Run("NVDS", testNVDSSelect)
	t.Run("Organizations", testOrganizationsSelect)
	t.Run("PackageInfos", testPackageInfosSelect)
	t.Run("Platforms", testPlatformsSelect)
	t.Run("References", testReferencesSelect)
	t.Run("ScanResults", testScanResultsSelect)
	t.Run("Tasks", testTasksSelect)
	t.Run("Users", testUsersSelect)
	t.Run("VulnInfos", testVulnInfosSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Auths", testAuthsUpdate)
	t.Run("Changelogs", testChangelogsUpdate)
	t.Run("Confidences", testConfidencesUpdate)
	t.Run("Containers", testContainersUpdate)
	t.Run("Cpes", testCpesUpdate)
	t.Run("CveDetails", testCveDetailsUpdate)
	t.Run("DistroAdvisories", testDistroAdvisoriesUpdate)
	t.Run("Groups", testGroupsUpdate)
	t.Run("JVNS", testJVNSUpdate)
	t.Run("NVDS", testNVDSUpdate)
	t.Run("Organizations", testOrganizationsUpdate)
	t.Run("PackageInfos", testPackageInfosUpdate)
	t.Run("Platforms", testPlatformsUpdate)
	t.Run("References", testReferencesUpdate)
	t.Run("ScanResults", testScanResultsUpdate)
	t.Run("Tasks", testTasksUpdate)
	t.Run("Users", testUsersUpdate)
	t.Run("VulnInfos", testVulnInfosUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Auths", testAuthsSliceUpdateAll)
	t.Run("Changelogs", testChangelogsSliceUpdateAll)
	t.Run("Confidences", testConfidencesSliceUpdateAll)
	t.Run("Containers", testContainersSliceUpdateAll)
	t.Run("Cpes", testCpesSliceUpdateAll)
	t.Run("CveDetails", testCveDetailsSliceUpdateAll)
	t.Run("DistroAdvisories", testDistroAdvisoriesSliceUpdateAll)
	t.Run("Groups", testGroupsSliceUpdateAll)
	t.Run("JVNS", testJVNSSliceUpdateAll)
	t.Run("NVDS", testNVDSSliceUpdateAll)
	t.Run("Organizations", testOrganizationsSliceUpdateAll)
	t.Run("PackageInfos", testPackageInfosSliceUpdateAll)
	t.Run("Platforms", testPlatformsSliceUpdateAll)
	t.Run("References", testReferencesSliceUpdateAll)
	t.Run("ScanResults", testScanResultsSliceUpdateAll)
	t.Run("Tasks", testTasksSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
	t.Run("VulnInfos", testVulnInfosSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("Auths", testAuthsUpsert)
	t.Run("Changelogs", testChangelogsUpsert)
	t.Run("Confidences", testConfidencesUpsert)
	t.Run("Containers", testContainersUpsert)
	t.Run("Cpes", testCpesUpsert)
	t.Run("CveDetails", testCveDetailsUpsert)
	t.Run("DistroAdvisories", testDistroAdvisoriesUpsert)
	t.Run("Groups", testGroupsUpsert)
	t.Run("JVNS", testJVNSUpsert)
	t.Run("NVDS", testNVDSUpsert)
	t.Run("Organizations", testOrganizationsUpsert)
	t.Run("PackageInfos", testPackageInfosUpsert)
	t.Run("Platforms", testPlatformsUpsert)
	t.Run("References", testReferencesUpsert)
	t.Run("ScanResults", testScanResultsUpsert)
	t.Run("Tasks", testTasksUpsert)
	t.Run("Users", testUsersUpsert)
	t.Run("VulnInfos", testVulnInfosUpsert)
}