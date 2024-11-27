# Solution for Receipt Processor Challenge
This is the Public Repository containing web service for Receipt Processor Challenge.

Here are few things I want you to consider before you move further ahead.

* Web Service used - RESTful web service

* Language selected - Go Language

* IDE used - Eclipse with GoClipse extension and Golang Development Tooling

## What are the different edge cases I came across?

1. **Invalid Payloads** - As mentioned, API endpoints accept only JSON payloads. rest all are considered invalid.
2. **Transfer protocol** - Because of using RESTful services, I was only able to use http and https protocols to successfully load API endpoints.
3. **id Value** - As mentioned previously, randomized UUID values were used based on the specification given by this problem. UUID is not tightly-coupled to the receipt once called, which means id's for the same receipt keep chaning.
4. **Input validation** - Input values for Retailer names and description are generally a mix of characters including special characters like @, !, # etc. Inputs accept all characters including alphanumeric and non-alphanumeric.
5. **Date and Time validation** - Errors like 13th month, 31st date, 25th hour of the day, 61st minute of the hour etc., are prone to happen as these entities do not exist.
6. **HTTP methods for endpoints** - Process Receipts endpoint accepts only POST whereas Get Points endpoint accepts only GET. Using methods other than what were assigned like using GET for Process Receipts and PUT for Get Points etc., would result in "Method not allowed" error. 

## Files in my project directory

1. **api.yml**  

   Description: yml file containing OpenAPI documentation.

   Directory: Receipt-Processor/config
   
2. **morning-receipt.json**
   
   Description: Example JSON payload.
   
   Directory: Receipt-Processor/examples 
    
3. **simple-receipt.json**

   Description: Example JSON payload.
   
   Directory: Receipt-Processor/examples 

4. **main.go**

   Description: Entry point of the Go application.

   Directory: Main folder

5. **receipt-handler.go**

    Description: Creates http handlers for both API endpoints.

    Directory: Receipt-Processor/handlers

6. **Receipt.go**

   Description: Determines the structure of the Receipt in JSON format. Contains Retailer, PurchaseDate, PurchaseTiem, Items, and Total.

   Directory: Receipt-Processor/models

7. **point_calculator.go**

   Description: Containes CalculatePoints function responsible for calculating points based on the given rules.

   Directory: Receipt-Processor/services

## Results:

1. For the first example test case i.e., morning-receipt.json, the output is,
      1. Process Receipts
         * Method: POST
         * Path: *http://localhost:8080/receipts/process*
         
         Response:
         
               {
                "id": "c3c55e53-259a-4ca6-aacd-68d3f9aa3493"
               }

      2. Get Points
         * Method: GET
         * Path: *http://localhost:8080/receipts/c3c55e53-259a-4ca6-aacd-68d3f9aa3493/points*
         
         Response:
         
               {
                "points": 15
               }

2. For the second example test case i.e., simple-receipt.json, the output is,
      1. Process Receipts
         * Method: POST
         * Path: *http://localhost:8080/receipts/process*
         
         Response:
         
               {
                "id": "2aae0161-b571-40e7-a2cd-e115349c6398"
               }

      2. Get Points
         * Method: GET
         * Path: *http://localhost:8080/receipts/2aae0161-b571-40e7-a2cd-e115349c6398/points*
         
         Response:
         
               {
                "points": 31
               } 
