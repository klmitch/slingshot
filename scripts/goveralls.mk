# Add support for goveralls
GOVERALLS = ./$(TOOLDIR)/goveralls
TOOLS     += github.com/mattn/goveralls

# Travis-specific target for submitting coverage to coveralls.io;
# explicitly undocumented
goveralls: $(COVER_OUT) $(GOVERALLS)
	$(GOVERALLS) -coverprofile=$(COVER_OUT) -service=travis-ci
