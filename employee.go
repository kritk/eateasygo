package hello
import (
        "time"

        "golang.org/x/net/context"

        "google.golang.org/appengine/datastore"
)

type Employee struct {
        FirstName          string
        LastName           string
        HireDate           time.Time
        AttendedHRTraining bool
}

func f(ctx context.Context) string{
        // ...
        employee := &Employee{
                FirstName: "Antonio",
                LastName:  "Salieri",
                HireDate:  time.Now(),
        }
        employee.AttendedHRTraining = true

        key := datastore.NewIncompleteKey(ctx, "Employee", nil)
        if _, err := datastore.Put(ctx, key, employee); err != nil {
                return err.Error()
        }else{
                return "Saved"
        }
        // ...
}