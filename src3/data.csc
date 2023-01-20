allow arr = [
    "stringobject",
    1,
    function() { return 1.0 },
    1.9
]

println("original array --> " + sprint(arr));
println("Reversed array --> " + sprint(reverse(arr)));

println("data")