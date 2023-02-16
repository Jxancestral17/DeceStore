<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>



<br />
<div align="center">
 
<h3 align="center">Decentralized Store</h3>

  <p align="center">
    For study purposes in go lang network and distributed services.
    <br />
  </p>
</div>




<!-- GETTING STARTED -->
## Getting Started

Clone the project & setup GOLang

  * Terminal
    ```
    go run main.go
    ```


  * Terminal
   ```
   curl -X POST localhost:8080 -d \
     '{"record": {"value": "TGV0J3MgR28gIzEK"}}'
   ```
  
 * Terminal
  ```
  curl -X POST localhost:8080 -d \
    '{"record": {"value": "TGV0J3MgR28gIzIK"}}'
  ```
  
  * Terminal
  ```
  curl -X POST localhost:8080 -d \
    '{"record": {"value": "TGV0J3MgR28gIzMK"}}'
  ```
  
  
   * Terminal
  ```
  curl -X GET localhost:8080 -d '{"offset": 0}'
  ```
  
  
   * Terminal
  ```
curl -X GET localhost:8080 -d '{"offset": 1}'
  ```
  
  
   * Terminal
  ```
 curl -X GET localhost:8080 -d '{"offset": 2}'
  ```
