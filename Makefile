clean:
	go clean -cache

.SILENT: test
test:
	go test -run TestGetLanguageByExtension ./internal/tests
	go test -run TestTryGetLanguageByExtension ./internal/tests
	go test -run TestGetLanguage ./internal/tests
	go test -run TestTryGetLanguage ./internal/tests

	go test -run TestGetPaste ./internal/tests
	go test -run TestTryGetPaste ./internal/tests

	go test -run TestCreatePaste ./internal/tests
	go test -run TestTryCreatePaste ./internal/tests

	go test -run TestCreatePrivatePaste ./internal/tests

	go test -run TestDeletePaste ./internal/tests
	go test -run TestTryDeletePaste ./internal/tests

	go test -run TestEditPaste ./internal/tests
	go test -run TestTryEditPaste ./internal/tests

	go test -run TestBulkDeletePastes ./internal/tests
	go test -run TestTryBulkDeletePastes ./internal/tests

	go test -run TestTimeExpiresInProperly ./internal/tests

	go test -run TestUserExists ./internal/tests
	go test -run TestTryGetUser ./internal/tests
	go test -run TestGetUser ./internal/tests
	go test -run TestGetSelfUser ./internal/tests
	go test -run TestTryGetSelfUser ./internal/tests
	go test -run TestGetSelfPastesByAmount ./internal/tests
	go test -run TestTryGetSelfPastesByAmount ./internal/tests
	go test -run TestGetSelfPasteIdsByAmount ./internal/tests
	go test -run TestTryGetSelfPasteIdsByAmount ./internal/tests
	go test -run TestGetSelfPasteIds ./internal/tests
	go test -run TestTryGetSelfPasteIds ./internal/tests
	go test -run TestGetSelfPastes ./internal/tests
	go test -run TestTryGetSelfPastes ./internal/tests