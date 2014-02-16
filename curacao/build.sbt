name := "curacao-benchmark"

organization := "com.kolich"

scalaVersion := "2.10.2"

version := "1.0"

resolvers ++= Seq(
  "markkolich.github.io repo" at "http://markkolich.github.io/repo"
)

libraryDependencies ++= Seq(
  "com.kolich.curacao" % "curacao" % "2.0-M10",
  "com.kolich.curacao" % "curacao-gson" % "2.0-M10",
  "org.eclipse.jetty" % "jetty-webapp" % "9.1.1.v20140108" % "container",
  "org.eclipse.jetty" % "jetty-plus" % "9.1.1.v20140108" % "container",
  "javax.servlet" % "javax.servlet-api" % "3.0.1" % "provided",
  "org.slf4j" % "slf4j-api" % "1.7.2" % "compile",
  "ch.qos.logback" % "logback-core" % "1.0.7" % "compile",
  "ch.qos.logback" % "logback-classic" % "1.0.7" % "compile"
)

classDirectory in Compile <<= baseDirectory(new File(_, "target/classes"))

seq(webSettings :_*)
