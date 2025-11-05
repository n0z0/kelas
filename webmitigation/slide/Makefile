# Makefile for Web Application Security Training Presentation

LATEX_COMPILER = pdflatex
LATEX_OPTIONS = -interaction=nonstopmode -shell-escape
MAIN_FILE = webapp_security_training.tex
OUTPUT_FILE = webapp_security_training.pdf

# Default target
all: $(OUTPUT_FILE)

# Compile the LaTeX document
$(OUTPUT_FILE): $(MAIN_FILE) *.tex
	$(LATEX_COMPILER) $(LATEX_OPTIONS) $(MAIN_FILE)
	$(LATEX_COMPILER) $(LATEX_OPTIONS) $(MAIN_FILE)
	$(LATEX_COMPILER) $(LATEX_OPTIONS) $(MAIN_FILE)

# Clean auxiliary files
clean:
	rm -f *.aux *.log *.toc *.out *.nav *.snm *.vrb *.fls *.fdb_latexmk

# Very clean (remove PDF as well)
veryclean: clean
	rm -f $(OUTPUT_FILE)

# View the PDF (requires a PDF viewer)
view: $(OUTPUT_FILE)
	xdg-open $(OUTPUT_FILE) || open $(OUTPUT_FILE) || start $(OUTPUT_FILE)

# Check for LaTeX errors
check: $(MAIN_FILE)
	$(LATEX_COMPILER) $(LATEX_OPTIONS) $(MAIN_FILE) | grep -E "(Error|Warning)"

# Help target
help:
	@echo "Available targets:"
	@echo "  all       - Compile the LaTeX document (default)"
	@echo "  clean     - Remove auxiliary files"
	@echo "  veryclean - Remove all generated files"
	@echo "  view      - Open the PDF in a viewer"
	@echo "  check     - Check for LaTeX errors"
	@echo "  help      - Show this help message"

.PHONY: all clean veryclean view check help