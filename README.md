<h1 align="center">Naive Bayes Spam Detection From Scratch <h1>
<p align="center">
 <a href="#license">
	<img src="https://img.shields.io/badge/License-MIT-blue?style=for-the-badge" alt="License"></a>
<a href="https://github.com/txmxthy/Spam-Detection-Naive-Bayes/issues">
	<img src="https://img.shields.io/github/issues/txmxthy/Spam-Detection-Naive-Bayes?style=for-the-badge" alt="issues - Spam-Detection-Naive-Bayes"></a>
<a href="https://github.com/txmxthy/Spam-Detection-Naive-Bayes">
	<img src="https://img.shields.io/github/stars/txmxthy/Spam-Detection-Naive-Bayes?style=for-the-badge" alt="stars - Spam-Detection-Naive-Bayes"></a>
<a href="https://github.com/txmxthy/Spam-Detection-Naive-Bayes">
	<img src="https://img.shields.io/github/forks/txmxthy/Spam-Detection-Naive-Bayes?style=for-the-badge" alt="forks - Spam-Detection-Naive-Bayes"></a>
</p>




<p align="center">
	<a href="https://golang.org">
		<img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" alt="Go"></a>
</p>


## About
This project is home to a naive bayesian classifier implemented from scratch in go. 
  
  This uses [Baye's Theorm](https://plato.stanford.edu/entries/bayes-theorem/) to calculate the probability that a piece of data belongs to a specific class based on prior knowledge. 
  In this case, the program takes two command-line arguments to a labelled and unlabelled dataset and applies the knowledge gained from the labelled dataset to label the other set, and display its findings.
  
  The dataset is comprised of 12 binary features and a binary class label. Feature selection is not part of this project.
  
  This project was mainly completed as a way to learn and practice go, it was not intended for practical or varied use; some functions are designed around the specific datasets. 
  
  Tested on Mac and Linux.


### Features

  Five Step Process
- Separate Data By Classes.
- Summarize Dataset.
- Summarize Data By Classes.
- Run Gaussian Probability Density Function.
- Estimate Class Probabilities.

 
  

## Usage

Build the program from source using
```zsh
 go build main.go
 ```
  
To run the program, first make sure it is marked as executable with
```zsh
  chmod +x main
 ```

  Then run the program with
```zsh
  ./main <path-to>spamLabelled.dat <path-to>spamUnlabelled.dat
  ```


  
## Contributing

Feel free to fork or make contributions. Any feedback is always welcome!

  
