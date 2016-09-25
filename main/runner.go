package main

//TODO: Ensure that the program is writting to stdout so that carriage returns work properly

/*******************************************************************************************/
/**                                                                                       **/
/**                                       Main                                            **/
/**                                                                                       **/
/*******************************************************************************************/

func main() {
	loadSettingsFile()
	loadPasswordFile()

	authenticate()

	loadRecords()
}
