<h1 align="center">Naive Bayes Spam Detection from Scratch <h1>
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
This project is home to a naive bayesian classifier implemented from scratch in go. The dataset this program was designed for is comprised of 12 binary features followed by binary class label, and each feature corresponds to indicators/attributes collected from spam and legitimate emails.
### Bayes Theorem
[Baye's Theorem](https://plato.stanford.edu/entries/bayes-theorem/) states that the probability of a hypothesis H conditional on a given body of data E is the ratio of the unconditional probability of the conjunction of the hypothesis with the data to the unconditional probability of the data alone. 
	
Baye's theorem is defined as the probability of H conditional on E is defined as PE(H) = P(H & E)/P(E), provided that both terms of this ratio exist and P(E) > 0.
  

  
### Application
Baye's theorem is used to calculate the conditional probability of an input belonging to a specific class based on prior knowledge. 
In this case, the program takes two command-line arguments to a labelled and unlabelled dataset. First, it applies the knowledge gained by the model's during the training phase from the labelled dataset to determine the conditional probability of a given feature belonging to the spam or non-spam classes. It then uses this probability to label the unlabelled set, before displaying its findings.
  

  
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

  
