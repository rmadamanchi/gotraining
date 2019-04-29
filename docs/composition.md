# Composition

## Background: Functions vs Methods -  Pointer vs Value Receivers
- Go has both functions and methods. In Go, a method is a function that is declared with a receiver. 
A receiver is a value or a pointer of a named or struct type.
All the methods for a given type belong to the type’s method set.
- You can treat the receiver as if it was an argument being passed to the method. 
All the same reasons why you might want to pass by value or pass by reference apply.

```
package main

import "fmt"

type User struct {
	name string
}

func (u User) SayName() {
	fmt.Println(u.name)
}

func (u User) ChangeName() {
	u.name = "Tom"
}

func main() {
	user := User {
		name: "Bob",
	}

	user.SayName()
	user.ChangeName()
	user.SayName()
}
```

Reasons why you would want to pass by reference as opposed to by value:
- You want to actually modify the receiver (“read/write” as opposed to just “read”)
- The struct is very large and a deep copy is expensive
- Consistency: if some of the methods on the struct have pointer receivers, the rest should too. This allows predictability of behavior

The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, 
but pointer methods can only be invoked on pointers.

## Type of relationships between objects
- Single Inheritance
- Multiple Inheritance
- Composition

### Composition vs Inheritance

#### Problems with traditional inheritance

- Easy to abuse and create complex type hierarchies
- Not able to inherit from multiple classes.
- Composition offers better testability (can mock composed types)
- Creates tight coupling. Hard to replace implementation of inheriting type with a better version

#### Examples of inheritance fail 
```
class Stack extends ArrayList {
    public void push(Object value) { … }
    public Object pop() { … }
}
```

```
Dog
  .poop()
   .bark()
  
Cat
  .poop()
  .meow()
```

```
Animal
  .poop()
    Dog
      .bark()
    Cat
      .meow()
```

```
Animal
  .poop()
    Dog
      .bark()
    Cat
     .meow()
    RobotDog
      .poop() - not implemented
      .bark()    
```


>  If you could do Java over again, what would you change?

>  I'd leave out classes - James Gosling, Inventor of Java

>   The real problem is't classes per se, but rather implementation inheritance
    (the extends relationship). Interface inheritance (the implements relationship)
    is preferable. You should avoid implementation inheritance whenever possible.

#### The fragile base-class problem
Base classes are considered fragile because you can modify a base class in a seemingly safe way,
but this new behavior, when inherited by the derived classes, might cause the derived classes
to malfunction.

The really big problem with inheritance is that you’re encouraged to predict the future.
Inheritance encourages you to build this taxonomy of objects very early on in your project,
and you are most likely going to make design mistakes doing that, because humans cannot
predict the future (even though it feels like we can), and getting out of these
inheritance taxonomies hard.

## Composition/Embedding in Go

- While Go does'nt have an "object" type, Go "struct" types contain fields and methods.  
- Go does not provide the typical, type-driven notion of subclassing, but it does have the ability to “borrow” 
pieces of an implementation by embedding types within a struct or interface.
- Instead of inheritance, Go strictly follows the composition over inheritance principle.

**Has-a relationship:**

```
package main

import "fmt"

type Person struct {
	Name string
	Address Address
}

type Address struct {
	Street string
	City   string
	State  string
}

func (p *Person) Speak() {
    fmt.Println("my name is: ", p.Name)
    fmt.Println("my address is: ",  p.Address.Street, p.Address.City, p.Address.State)
}

func main() {
	p := NewPerson()

	p.Speak()
}

func NewPerson() *Person {
	return &Person{
		Name: "Steve",
		Address: Address{
			Street: "123 Main St",
			City:   "Austin",
			State:  "TX",
		},
	}
}

```

**Is-a relationship (kinda)**

Embedded types
By embedding the structs directly, the methods of embedded types come along for free. 
There's an important way in which embedding differs from subclassing. 
When we embed a type, the methods of that type become methods of the outer type, but when they are invoked the receiver of the method is the inner type, not the outer one.

```
package main

import "fmt"

type Person struct {
	Name string
	Address Address
}

type Address struct {
	Street string
	City   string
	State  string
}

func (person *Person) Speak() {
	fmt.Println("my name is", person.Name)
	fmt.Println("my address is",  person.Address.Street, person.Address.City, person.Address.State)
}

type Worker struct {
	Person
	Profession string
}

func main() {
	worker := NewWorker()

	worker.Speak()
}

func NewWorker() *Worker {
	worker := &Worker{}
	worker.Name = "Steve"
	worker.Address = Address{
		Street: "123 Main St",
		City:   "Austin",
		State:  "TX",
	}
	worker.Profession = "Engineer"

	return worker
}

```

```
func (worker *Worker) Speak() {
	fmt.Println("my name is", worker.Name, "and i'm a" + worker.Profession)
	fmt.Println("my address is",  worker.Address.Street, worker.Address.City, worker.Address.State)
}
```

Worker cannot be used in place of Person. There is an embedded Person in Worker.

You can't pass a Worker to a function that is looking for a Person
```
func main() {
	worker := NewWorker()

	worker.Speak()

	NeedAPerson(worker)
}

func NeedAPerson(person *Person) {
	fmt.Println("I got", person.Name)
}

```

```

type Student struct {
	Person
	Grade string
}

func (student *Student) Speak() {
	fmt.Println("my name is", student.Name, "and i'm in grade", student.Grade)
	fmt.Println("my address is",  student.Address.Street, student.Address.City, student.Address.State)
}

func main() {
	pepole := []Person{
		Worker{},
		Student{},
	}
}

```

**Ambiguous Methods**


```
type StudentWorker struct {
	Worker
	Student
}

func NewStudentWorker() *StudentWorker {
	return &StudentWorker{
		Worker: Worker{
			Person: Person {
				Name: "Tom",
			},
			Profession: "Engineer",
		},
		Student: Student{
			Person: Person {
				Name: "Tom",
			},
			Grade: "10",
		},
	}
}
func main() {
	sw := NewStudentWorker()

	sw.Speak()
}
```

```
func (studentWorker *StudentWorker) Speak() {
	fmt.Println("my name is", studentWorker.Name, "and i'm in grade", studentWorker.Grade, "and i'm a", studentWorker.Profession)
	fmt.Println("my address is",  studentWorker.Address.Street, studentWorker.Address.City, studentWorker.Address.State)
}
```

- The Person type is not adding any real value. and making the code less readable, simple or adaptable
- Traditional issues with inheritance



**Using Interfaces**


```
package main

import "fmt"

type Speaker interface {
	Speak()
}

type Address struct {
	Street string
	City   string
	State  string
}

type Worker struct {
	Name string
	Address Address
	Profession string
}

func (worker *Worker) Speak() {
	fmt.Println("my name is", worker.Name, "and i'm a",  worker.Profession)
	fmt.Println("my address is",  worker.Address.Street, worker.Address.City, worker.Address.State)
}

func NewWorker() *Worker {
	return &Worker {
		Name: "Steve",
		Profession: "Engineer",
		Address: Address{
			Street: "123 Main St",
			City:   "Austin",
			State:  "TX",
		},
	}
}

type Student  struct {
	Name string
	Address Address
	Grade string
}

func (student *Student) Speak() {
	fmt.Println("my name is", student.Name, "and i'm in grade",  student.Grade)
	fmt.Println("my address is",  student.Address.Street, student.Address.City, student.Address.State)
}

func NewStudent() *Student {
	return &Student {
		Name: "Bob",
		Grade: "12",
		Address: Address{
			Street: "123 Main St",
			City:   "Austin",
			State:  "TX",
		},
	}
}

func NeedASpeaker(speaker Speaker) {
	speaker.Speak()
}

func NeedSpeakers(speakers []Speaker) {
	fmt.Println("----")
	for _, speaker := range speakers {
		speaker.Speak()
		fmt.Println("----")
	}
}

func main() {
	student := NewStudent()
	NeedASpeaker(student)

	worker := NewWorker()
	NeedASpeaker(worker)

	people := []Speaker {
		student, worker,
	}
	NeedSpeakers(people)
}
```


- Try to avoid these type hierarchies by thinking about the idea of common behavior than the idea of common state
- Smaller interfaces with specific behaviors
- Idiomatic to end single method interface names with `er`

```
package main

type PersonInterface interface {
	Speak() string
}

type ValuePerson struct {}

func (vp ValuePerson) Speak() string {
	return "hey"
}

type PointerPerson struct {}

func (pp *PointerPerson) Speak() string {
	return "hey"
}

func AskToSpeak(pi PersonInterface) {
	pi.Speak()
}

func main() {
	vp := ValuePerson{}
	AskToSpeak(vp)
	AskToSpeak(&vp)

	pp := PointerPerson{}
	AskToSpeak(pp)
	AskToSpeak(&pp)
}
```
The method set of the corresponding pointer type *T is the set of all methods with receiver *T or T
The method set of any other type T consists of all methods with receiver type T.
=> The method set of the corresponding type T does not consists of any methods with receiver type *T.


### Embedding interfaces in Interfaces and structs

```
package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReaderWriter interface {
	Reader
	Writer
}

type Editor struct {

}

func (e Editor) Read() {
	fmt.Println("Reading")
}

func (e Editor) Write() {
	fmt.Println("Writing")
}

func PerfomRead(r Reader) {
	r.Read()
}

func PerformWrite(w Writer) {
	w.Write()
}

func ReadAndWrite(rw ReaderWriter) {
	rw.Read()
	rw.Write()
}


type BetterReader struct {
	Reader
}

func main() {
	PerfomRead(Editor{})
	PerformWrite(Editor{})

	ReadAndWrite(Editor{})

	br := BetterReader{Reader: Editor{}}
	PerfomRead(br)
}

```


-  Declare the set of behaviors as discrete interface types first. Then think about how they can be composed into a larger set of behaviors.
-  Make sure each function or method is very specific about the interface types they accept. Only accept interface types for the behavior you are using in that function or method. This will help dictate the larger interface types that are required.
-  Think about embedding as an inner and outer type relationship. Remember that through inner type promotion, everything that is declared in the inner type is promoted to the outer type. However, the inner type value exists in and of itself as is always accessible based on the rules for exporting.
- Type embedding is not subtyping nor subclassing. Concrete type values represent a single type and can’t be assigned based on any embedded relationship.
- The compiler can arrange interface conversions between related interface values. Interface conversion, at compile time, doesn’t care about concrete types - it knows what to do merely based on the interface types themselves, not the implementing concrete values they could contain.


## Interface Pollution

```
package tcp
// Server defines a contract for tcp servers.
type Server interface {
    Start() error
    Stop() error
    Wait() error
}

// server is our Server implementation.
type server struct {
    /* impl */
}

// NewServer returns an interface value of type Server
// with an xServer implementation.
func NewServer(host string) Server {
    return &server{host}
}

// Start allows the server to begin to accept requests.
func (s *server) Start() error {
    /* impl */
}

// Stop shuts the server down.
func (s *server) Stop() error {
    /* impl */
}

// Wait prevents the server from accepting new connections.
func (s *server) Wait() error {
    /* impl */
}
```

vs

```
type Server struct {
    /* impl */
}

# Have the NewServer function return a pointer of the concrete type instead of the interface type.

func NewServer(host string) *Server {
	return &Server{host}
}
```

- Since implementing types don't have to declare interfaces, interfaces can be designed after implementation.
- Do you need to declare interfaces for testing/mocking purposes?
- Mocking struct types using self-declaring interfaces

Use an interface:
- When users of the API need to provide an implementation detail.
- When API’s have multiple implementations they need to maintain internally.
- When parts of the API that can change have been identified and require decoupling.

Don't use an interface:
- for the sake of using an interface.
- to generalize an algorithm.
- when users can declare their own interfaces.
- if it's not clear how the interface makes the code better.