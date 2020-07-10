mkdir classes test-classes
javac -d classes src/Pair.java
javac -d test-classes -classpath classes:../../lib/junit4/junit-4.13.jar:../../lib/junit4/hamcrest-all-1.3.jar test/PairTest.java
