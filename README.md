# go_otp_verify

# Go OTP Verification using Twilio

This is an example of using Twilio to send OTP verification code to a user and verifying the code.

## Setup

1. Clone the repo
2. Create a Twilio account and get your account SID and auth token
3. Create a Twilio verify service and get the service SID
4. Create a .env file in the root of the project and add the in the above credentials from Twilio

```env
TWILIO_ACCOUNT_SID=<ACCOUNT SID>
TWILIO_AUTHTOKEN=<AUTH TOKEN>
TWILIO_SERVICES_ID=<SERVICE ID>
```

## API Documentation

Send OTP
Send a POST request to the /otp endpoint with the following body to send an OTP to a user's phone number
POST /otp
Request Body

    ```json
    {
        "phone_number": "+2348123456789"
    }
    ```
    ```bash curl -H "Content-Type: application/json" -X POST -d '{"phoneNumber": "+"}' http://localhost:8000/otp```

Be sure to include the country code in the phone number

## Verify OTP

Verify a user's OTP by sending a POST request to the /verify endpoint with the following body that contains the phone number and the OTP code received by the user

verifyOTP
Request Body
