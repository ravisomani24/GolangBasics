#### **Packages are used to organize Go source code for better reusability and readability. Packages are a collection of Go sources files that reside in the same directory. Packages provide code compartmentalization and hence it becomes easy to maintain Go projects.**

### **Usage**:

**package packagename** specifies that a particular source file belongs to package packagename. This should be the first line of every go source file.

### **Go Module**:

**A Go Module is nothing but a collection of Go packages.**
Now this question might come to your mind. Why do we need Go modules to create a custom package? The answer is the import path for the custom package we create is derived from the name of the go module.
In addition to this, all the other third-party packages(such as source code from github) which our application uses will be present in the go.mod file along with the version. This go.mod file is created when we create a new module.

### **Example**:

1. **cd learnpackages**: Make sure we are in the folder of which we are creating module
2. **go mod init learnpackages**: This will create a go.mod file. Because of this our code will understand where is the module learnpackages is and then only can find packages inside this module
3. **import packages**: Now we can import custom packages from module learnpackages

### **init function**:
Each package in Go can contain an init function. The init function must not have any return type and it must not have any parameters. The init function cannot be called explicitly in our source code.
It will be called automatically when the package is initialized. The init function has the following syntax

`func init() {
}`

The init function can be used to perform initialization tasks and can also be used to verify the correctness of the program before the execution starts.

The order of initialization of a package is as follows

1. Package level variables are initialised first
2. init function is called next. A package can have multiple init functions (either in a single file or distributed across multiple files) and they are called in the order in which they are presented to the compiler.

If a package imports other packages, the imported packages are initialized first.
A package will be initialized only once even if it is imported from multiple packages.

Sometimes we need to import a package just to make sure the initialization takes place even though we do not need to use any function or variable from the package.
For example, we might need to ensure that the init function of the simpleinterest package is called even though we plan not to use that package anywhere in our code. The _ blank identifier can be used in this case too as shown below.