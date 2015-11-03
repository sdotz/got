# Go Training

Interface satisfaction is at the value level based on the rules for method sets:
You can't call an interface method that is implemented with a pointer receiver by passing an object by value. The object might not exist in memory. Pass a pointer instead. If the inteface method is implemented with a value receiver, you may pass a value or pointer.