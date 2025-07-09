#!/usr/bin/env python3

"""
PDF Generator for Nomado Gantt Chart
This script converts the HTML Gantt chart to PDF format
"""

import subprocess
import sys
import os

def check_and_install_requirements():
    """Check if required packages are installed, install if needed"""
    try:
        import weasyprint
        print("✓ WeasyPrint is available")
        return True
    except ImportError:
        print("Installing WeasyPrint...")
        try:
            subprocess.check_call([sys.executable, "-m", "pip", "install", "weasyprint"])
            import weasyprint
            print("✓ WeasyPrint installed successfully")
            return True
        except Exception as e:
            print(f"❌ Failed to install WeasyPrint: {e}")
            return False

def convert_html_to_pdf():
    """Convert HTML file to PDF"""
    try:
        import weasyprint
        
        # Define file paths
        html_file = "/home/diesel/Desktop/nomado-houses/GANTT_CHART.html"
        pdf_file = "/home/diesel/Desktop/nomado-houses/GANTT_CHART.pdf"
        
        # Check if HTML file exists
        if not os.path.exists(html_file):
            print(f"❌ HTML file not found: {html_file}")
            return False
        
        print(f"Converting {html_file} to PDF...")
        
        # Create PDF from HTML
        html_doc = weasyprint.HTML(filename=html_file)
        html_doc.write_pdf(pdf_file)
        
        print(f"✓ PDF created successfully: {pdf_file}")
        return True
        
    except Exception as e:
        print(f"❌ Error converting to PDF: {e}")
        return False

def main():
    print("🚀 Nomado Gantt Chart PDF Generator")
    print("=" * 50)
    
    # Check and install requirements
    if not check_and_install_requirements():
        print("\n💡 Alternative: You can open GANTT_CHART.html in your browser and print to PDF manually")
        return
    
    # Convert HTML to PDF
    if convert_html_to_pdf():
        print("\n✅ PDF generation completed successfully!")
        print("📄 Your Gantt chart PDF is ready at: GANTT_CHART.pdf")
    else:
        print("\n❌ PDF generation failed")
        print("💡 Alternative: You can open GANTT_CHART.html in your browser and print to PDF manually")

if __name__ == "__main__":
    main()
