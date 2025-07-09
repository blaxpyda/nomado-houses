#!/bin/bash

# Nomado Gantt Chart PDF Generator Script
# This script helps you convert the HTML Gantt chart to PDF

echo "🚀 Nomado Gantt Chart PDF Generator"
echo "=================================="
echo ""

# Check if the HTML file exists
if [ ! -f "GANTT_CHART.html" ]; then
    echo "❌ GANTT_CHART.html not found in current directory"
    exit 1
fi

echo "📋 HTML Gantt chart found: GANTT_CHART.html"
echo ""

# Check for available PDF conversion tools
if command -v wkhtmltopdf &> /dev/null; then
    echo "✓ wkhtmltopdf found. Converting to PDF..."
    wkhtmltopdf --page-size A4 --margin-top 0.75in --margin-right 0.75in --margin-bottom 0.75in --margin-left 0.75in --print-media-type GANTT_CHART.html GANTT_CHART.pdf
    echo "✅ PDF created: GANTT_CHART.pdf"
    
elif command -v chromium-browser &> /dev/null; then
    echo "✓ Chromium found. Converting to PDF..."
    chromium-browser --headless --disable-gpu --print-to-pdf=GANTT_CHART.pdf --virtual-time-budget=1000 file://$(pwd)/GANTT_CHART.html
    echo "✅ PDF created: GANTT_CHART.pdf"
    
elif command -v google-chrome &> /dev/null; then
    echo "✓ Google Chrome found. Converting to PDF..."
    google-chrome --headless --disable-gpu --print-to-pdf=GANTT_CHART.pdf --virtual-time-budget=1000 file://$(pwd)/GANTT_CHART.html
    echo "✅ PDF created: GANTT_CHART.pdf"
    
elif command -v firefox &> /dev/null; then
    echo "💡 Firefox found. Manual conversion instructions:"
    echo ""
    echo "1. Open Firefox"
    echo "2. Go to: file://$(pwd)/GANTT_CHART.html"
    echo "3. Press Ctrl+P (or Cmd+P on Mac)"
    echo "4. Choose 'Save as PDF' as destination"
    echo "5. Click 'Save'"
    echo ""
    # Try to open the file in Firefox
    firefox "file://$(pwd)/GANTT_CHART.html" &
    
else
    echo "💡 No automatic PDF conversion tool found."
    echo ""
    echo "📖 Manual Conversion Instructions:"
    echo "================================="
    echo ""
    echo "Option 1 - Using any web browser:"
    echo "1. Open your web browser (Chrome, Firefox, Safari, Edge)"
    echo "2. Open the file: file://$(pwd)/GANTT_CHART.html"
    echo "3. Press Ctrl+P (or Cmd+P on Mac) to open print dialog"
    echo "4. Choose 'Save as PDF' or 'Microsoft Print to PDF'"
    echo "5. Adjust settings:"
    echo "   - Paper size: A4"
    echo "   - Margins: Normal"
    echo "   - Include background graphics: Yes"
    echo "6. Save as GANTT_CHART.pdf"
    echo ""
    echo "Option 2 - Install wkhtmltopdf:"
    echo "sudo apt install wkhtmltopdf"
    echo "Then run this script again."
    echo ""
fi

echo ""
echo "📁 Files created:"
echo "  - GANTT_CHART.md (Markdown version)"
echo "  - GANTT_CHART.html (HTML version - ready for printing)"
if [ -f "GANTT_CHART.pdf" ]; then
    echo "  - GANTT_CHART.pdf (PDF version)"
fi
echo ""
echo "✨ Your Nomado development roadmap is ready!"
