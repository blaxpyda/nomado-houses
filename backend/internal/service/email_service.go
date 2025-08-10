package service

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

// EmailService interface defines methods for email operations
type EmailService interface {
	SendVerificationEmail(email, firstName, verificationCode string) error
	SendWelcomeEmail(email, firstName string) error
	GenerateVerificationCode() string
}

// emailService implements EmailService
type emailService struct {
	smtpHost string
	smtpPort string
	username string
	password string
	fromName string
}

// NewEmailService creates a new email service
func NewEmailService() EmailService {
	return &emailService{
		smtpHost: os.Getenv("SMTP_HOST"),
		smtpPort: os.Getenv("SMTP_PORT"),
		username: os.Getenv("SMTP_USERNAME"),
		password: os.Getenv("SMTP_PASSWORD"),
		fromName: os.Getenv("SMTP_FROM_NAME"),
	}
}

// GenerateVerificationCode generates a random verification code
func (s *emailService) GenerateVerificationCode() string {
	bytes := make([]byte, 3) // 6 character hex code
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// SendVerificationEmail sends a verification email
func (s *emailService) SendVerificationEmail(email, firstName, verificationCode string) error {
	subject := "Verify Your Nomado Account"

	// HTML template for verification email
	htmlTemplate := `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Verify Your Email</title>
		<style>
			body { font-family: Arial, sans-serif; line-height: 1.6; margin: 0; padding: 0; background-color: #f4f4f4; }
			.container { max-width: 600px; margin: 0 auto; background: white; padding: 20px; border-radius: 10px; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
			.header { background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); color: white; padding: 30px; text-align: center; border-radius: 10px 10px 0 0; }
			.content { padding: 30px; }
			.verification-code { background: #f8f9ff; border: 2px dashed #667eea; padding: 20px; margin: 20px 0; text-align: center; border-radius: 8px; }
			.code { font-size: 32px; font-weight: bold; color: #667eea; letter-spacing: 4px; }
			.button { display: inline-block; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); color: white; padding: 12px 30px; text-decoration: none; border-radius: 25px; margin: 20px 0; }
			.footer { text-align: center; color: #666; font-size: 12px; margin-top: 30px; }
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<h1>🌍 Welcome to Nomado!</h1>
				<p>Your Gateway to Africa & Beyond</p>
			</div>
			<div class="content">
				<h2>Hello {{.FirstName}}!</h2>
				<p>Thank you for joining Nomado! To complete your registration and start your journey with us, please verify your email address.</p>
				
				<div class="verification-code">
					<p><strong>Your Verification Code:</strong></p>
					<div class="code">{{.VerificationCode}}</div>
				</div>
				
				<p>Enter this code on the verification page, or click the button below:</p>
				<div style="text-align: center;">
					<a href="{{.VerificationURL}}" class="button">Verify My Email</a>
				</div>
				
				<p><strong>Important:</strong> This code will expire in 24 hours for security reasons.</p>
				
				<hr style="border: 1px solid #eee; margin: 30px 0;">
				
				<h3>What's Next?</h3>
				<ul>
					<li>🏨 Book amazing accommodations across Africa</li>
					<li>✈️ Get help with visa applications</li>
					<li>🚌 Find the best transportation options</li>
					<li>💕 Discover romantic getaways with Nomado Love</li>
					<li>🎓 Plan educational trips with Little Nomads</li>
				</ul>
			</div>
			<div class="footer">
				<p>If you didn't create this account, please ignore this email.</p>
				<p>© 2025 Nomado. All rights reserved.</p>
				<p>Made with ❤️ for African travelers</p>
			</div>
		</div>
	</body>
	</html>`

	tmpl, err := template.New("verification").Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse email template: %w", err)
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, struct {
		FirstName        string
		VerificationCode string
		VerificationURL  string
	}{
		FirstName:        firstName,
		VerificationCode: verificationCode,
		VerificationURL:  fmt.Sprintf("%s/verify-email?email=%s&code=%s", os.Getenv("FRONTEND_URL"), email, verificationCode),
	})

	if err != nil {
		return fmt.Errorf("failed to execute email template: %w", err)
	}

	return s.sendEmail(email, subject, body.String(), true)
}

// SendWelcomeEmail sends a welcome email after verification
func (s *emailService) SendWelcomeEmail(email, firstName string) error {
	subject := "Welcome to Nomado - Let's Start Your Journey!"

	htmlTemplate := `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Welcome to Nomado</title>
		<style>
			body { font-family: Arial, sans-serif; line-height: 1.6; margin: 0; padding: 0; background-color: #f4f4f4; }
			.container { max-width: 600px; margin: 0 auto; background: white; padding: 20px; border-radius: 10px; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
			.header { background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); color: white; padding: 30px; text-align: center; border-radius: 10px 10px 0 0; }
			.content { padding: 30px; }
			.button { display: inline-block; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); color: white; padding: 12px 30px; text-decoration: none; border-radius: 25px; margin: 20px 10px; }
			.services { display: grid; grid-template-columns: 1fr 1fr; gap: 20px; margin: 20px 0; }
			.service { background: #f8f9ff; padding: 15px; border-radius: 8px; text-align: center; }
			.footer { text-align: center; color: #666; font-size: 12px; margin-top: 30px; }
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<h1>🎉 Welcome to Nomado, {{.FirstName}}!</h1>
				<p>Your verified account is ready to go!</p>
			</div>
			<div class="content">
				<h2>You're All Set! 🚀</h2>
				<p>Your email has been successfully verified and your Nomado account is now active. We're excited to help you explore Africa and beyond!</p>
				
				<div style="text-align: center; margin: 30px 0;">
					<a href="{{.DashboardURL}}" class="button">Start Exploring</a>
					<a href="{{.ServicesURL}}" class="button">View Services</a>
				</div>
				
				<h3>Popular Services to Get You Started:</h3>
				<div class="services">
					<div class="service">
						<h4>🏨 Hotels & Stays</h4>
						<p>Find perfect accommodations</p>
					</div>
					<div class="service">
						<h4>📋 Visa Help</h4>
						<p>Professional assistance</p>
					</div>
					<div class="service">
						<h4>🚌 Transportation</h4>
						<p>Flights, buses, car rentals</p>
					</div>
					<div class="service">
						<h4>💕 Special Experiences</h4>
						<p>Romance, family, luxury</p>
					</div>
				</div>
				
				<p>Need help? Our support team is available 24/7 to assist you with your travel plans.</p>
			</div>
			<div class="footer">
				<p>Follow us on social media for travel tips and deals!</p>
				<p>© 2025 Nomado. All rights reserved.</p>
			</div>
		</div>
	</body>
	</html>`

	tmpl, err := template.New("welcome").Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse welcome email template: %w", err)
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, struct {
		FirstName    string
		DashboardURL string
		ServicesURL  string
	}{
		FirstName:    firstName,
		DashboardURL: os.Getenv("FRONTEND_URL") + "/dashboard",
		ServicesURL:  os.Getenv("FRONTEND_URL") + "/services",
	})

	if err != nil {
		return fmt.Errorf("failed to execute welcome email template: %w", err)
	}

	return s.sendEmail(email, subject, body.String(), true)
}

// sendEmail sends an email using SMTP
func (s *emailService) sendEmail(to, subject, body string, isHTML bool) error {
	// Validate configuration
	if s.smtpHost == "" || s.smtpPort == "" || s.username == "" || s.password == "" {
		return fmt.Errorf("incomplete SMTP configuration")
	}

	from := s.username
	if s.fromName != "" {
		from = fmt.Sprintf("%s <%s>", s.fromName, s.username)
	}

	// Create message
	var message bytes.Buffer
	message.WriteString(fmt.Sprintf("From: %s\r\n", from))
	message.WriteString(fmt.Sprintf("To: %s\r\n", to))
	message.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	message.WriteString("MIME-Version: 1.0\r\n")

	if isHTML {
		message.WriteString("Content-Type: text/html; charset=UTF-8\r\n")
	} else {
		message.WriteString("Content-Type: text/plain; charset=UTF-8\r\n")
	}

	message.WriteString("\r\n")
	message.WriteString(body)

	// Set up authentication
	auth := smtp.PlainAuth("", s.username, s.password, s.smtpHost)

	// Send email
	addr := fmt.Sprintf("%s:%s", s.smtpHost, s.smtpPort)
	err := smtp.SendMail(addr, auth, s.username, []string{to}, message.Bytes())
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
