[
  {fruits: "apple",  sports: "tennis",       movie: "matrix"},
  {fruits: "banana", sports: "table tennis", movie: "matrix"},
  {fruits: "berry",  sports: "baseball",     movie: "not interested"},
  {fruits: "apple",  sports: "snowboarding", movie: "oceans11"}
].each do |items|
  File.open "sample.ltsv", "a" do |file|
    line = items.map{|k, v| "#{k}:#{v}" }.join "\t"
    file.puts line
  end
end
