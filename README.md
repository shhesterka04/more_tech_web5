# Documentation
## Go API Handlers
1. **GetOneBranch**
   _Endpoint: /api/branch/{branchId}_

   Method: GET

   Description: Retrieves information about the branch based on the provided branch ID. It also fetches forecasting data for the branch's client numbers.

   Parameters:

   - branchId (path): The ID of the branch.
   
   **Response**: Returns data about a branch including current client numbers and forecasted numbers for the week in the form of a table.


2. **GetBranchesByFilter**

   _Endpoint: /api/branches_

   Method: GET

   Description: Retrieves information about branches or ATMs based on provided filters.

   Query Parameters:
   - isOffice: Indicates if the filter is for office (1) or ATM (0).
   - qr: Indicates if the ATM supports QR (1 if yes).
   - nfc: Indicates if the ATM supports NFC (1 if yes).
   - blind: Indicates if the ATM supports facilities for the blind (1 if yes).
   - wheelchair: Indicates if the ATM supports wheelchair access (1 if yes).
   - face: Unknown filter.
   - allday: Indicates if the ATM is available all day (1 if yes).
   - officetype: Indicates type of office (1 for a specific type).
   
   **Response**: Returns a list of branches or ATMs that match the filter criteria.


3. **GetRecomBranch**

   _Endpoint: /api/branches/recommended_

   Method: POST

   Description: Recommends a branch based on starting and ending coordinates and the transport type provided in the request body.

   Request Body:

   - start: Coordinates of the starting location (latitude, longitude).
   - end: Coordinates of the ending location (latitude, longitude).
   - transportType: Type of transportation (e.g., car, bike, etc.).
   
   **Response**: Returns the recommended route.


##   Python API Handlers
   **Prophet Forecasting Handler**

   Description: This handler is responsible for:

   Reading a file named loadperweek.txt.

   Transforming the data and converting it to two dataframes: df1 for individuals and df2 for legal entities.

   Using the Prophet library to forecast future data for both individuals and legal entities for the next 7 days.
   Writing the forecast data to a file named forecast_data.txt.

   **Output**: Writes forecasted data into a file.
   
# Instructions for Running the Code


## Go:
###  Setup:

```
go get github.com/gorilla/mux
go get github.com/gorilla/handlers

go run main.go
```

Access:

Once the server is running, you can access the API using a browser or tools like curl or Postman by making requests to http://localhost:8080.

## Python:
### Setup:

```
pip install -r requirements.txt
cd src
uvicorn main:app --reload
```

