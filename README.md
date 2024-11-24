<!-- Back to Top Link-->
<a name="readme-top"></a>


<br />
<div align="center">
  <h1 align="center">Wikirace Solver using combination of BFS and IDS</h1>

  <p align="center">
    <h3>Back End</h3>
    <h4>Tugas Besar 2 IF2211 Strategi Algoritma</h4>
    <h3><a href="https://github.com/NoHaitch/Tubes2_BE_Chibye">Front End</a> & <a href="https://github.com/NoHaitch/Tubes2_BE_Chibye">Back End</a></h3>
    <br/>
    <a href="https://github.com/NoHaitch/Tubes2_BE_Chibye/issues">Report Bug</a>
    ·
    <a href="https://github.com/NoHaitch/Tubes2_BE_Chibye/issues">Request Feature</a>
<br>
<br>

[![Apache v2.0 License][license-shield]][license-url]

  </p>
</div>

<!-- CONTRIBUTOR -->
<div align="center" id="contributor">
  <strong>
    <h3>Made By:</h3>
    <h3>Kelompok Chibye</h3>
    <table align="center">
      <tr>
        <td>NIM</td>
        <td>Nama</td>
      </tr>
      <tr>
        <td>13522029</td>
        <td>Ignatius Jhon Hezkiel Chan</td>
      </tr>
      <tr>
        <td>13522091</td>
        <td>Raden Francisco Trianto Bratadiningrat</td>
      </tr>
      <tr>
        <td>13522098</td>
        <td>Suthasoma Mahardhika Munthe</td>
      </tr>
    </table>
  </strong>
  <br>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started-back-end">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>

## External Links

- [Spesifikasi](https://docs.google.com/document/d/1h6WY_NxfCBPrKkS84Crm2qAhrRA8DatL/edit)
- [QNA](https://docs.google.com/spreadsheets/d/1egeULRNv3ZrCrRexrbi7G4GkKwi_9KGasFIPAnhODfw/edit#gid=982607851)
- [Teams](https://docs.google.com/spreadsheets/d/14wDe_K5LjHpsEnQSoLrB4mNf98zTTP-0xWkXqoWDOMw/edit#gid=0)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ABOUT THE PROJECT -->
## About The Project

  For Tugas Besar 2, we are to make a solver for [Wikirace Game](https://en.wikipedia.org/wiki/Wikipedia:Wiki_Game). In summary, Wikirace is a game of finding the fastest way to get from a source Wikipedia page to a target page, where the number of links needed matters to the amount of time to reach the target page.

To solve Wikirace, we use two search algorithms, BFS and IDS. For learning reasons, we are not to use Wikipedia API but instead need to web scrape all the links in a Wikipedia Page. This causes us many problems regarding limited requests to Wikipedia.

Our Project is divided into a Front-end and a Back-end. Here are the links both repository:  
- Front-end: https://github.com/NoHaitch/Tubes2_FE_Chibye 
- Back-end: https://github.com/NoHaitch/Tubes2_BE_Chibye   


<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started Back-end

### Prerequisites

Project dependencies  

* Golang  
  You can find how to install golang here: https://go.dev/doc/install 

Golang library used:
- [gocolly](https://go-colly.org/)
- [gin](https://gin-gonic.com/)


<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Installation

How to install and use this project (Without Docker)

1. Clone the repo
   ```sh
   git clone https://github.com/NoHaitch/Tubes2_BE_Chibye
   ```
2. Go to src Directory
   ```sh
   cd src
   ```
3. Get all dependencies
   ```sh
   go get
   ```
4. Run the API
   ```sh
   go run .
   ``` 
<br>

How to install and use this project (With docker)
1. Create New Folder as the Parent for FrontEnd and BackEnd folder. Let's say Tubes2_Chibye
2. Clone the FrontEnd repository to the Tubes2_Chibye folder.
   ```sh
   git clone https://github.com/NoHaitch/Tubes2_FE_Chibye
   ```
3. Clone the BackEnd repository to the Tubes2_Chibye folder.
   ```sh
   git clone https://github.com/NoHaitch/Tubes2_BE_Chibye
   ```
4. Move the `docker-compose.yml` file in the FrontEnd folder, out to the Tubes2_Chibye folder so the file will be the same directory as the FrontEnd and BackEnd folder.
5. Open your terminal at the Tubes2_Chibye directory, and then type this command.
   ```sh
   docker compose up --build
   ```
6. If you want to run the web app again and had done the fifth step, just type docker compose up without the build tag
    ```sh
    docker compose up
    ```
7. Now, the web app is running with the frontend and backend functionality. Visit `http://localhost:3000/` to open the WikiRace web app.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- FEATURES -->
## Features

For Front-end this project uses React  
For Web Framework this project uses Gin   
For scrapping this project uses gocolly  

### 1. BFS Search
### 2. IDS Search
### 3. Suggestion for Search Bar
### 4. Validation for Title Input
### 5. Resulting Search shown using graph
### 6. Shows how many pages visited
### 7. Shows time taken to find the solution
### 8. Shows how many pages needed to reach the target
### 9. Caching for IDS search
### 10. Concurrency for IDS and BFS
### 11. Docker support

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTRIBUTING -->
## Contributing

If you want to contribute or further develop the program, please fork this repository using the branch feature.  
Pull Request is **permited and warmly welcomed**

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## Licensing

The code in this project is licensed under Apache 2.0 license.  

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<br>
<h3 align="center"> THANK YOU! </h3>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[issues-url]: https://github.com/NoHaitch/Tubes2_BE_Chibye/issues
[license-shield]: https://img.shields.io/badge/License-Apache--2.0_license-yellow
[license-url]: https://github.com/NoHaitch/Tubes2_BE_Chibye/blob/main/LICENSE
