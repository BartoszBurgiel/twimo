// Import core functionalities from react & react-native library
import React from "react";
import {
  View,
  Text,
  Button,
  TextInput,
  TouchableOpacity,
  StyleSheet
} from "react-native";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";

//TODO: Implement real credential check!

// Define HomeScreen
const HomeScreen = ({ navigation }) => {
  return (
    <View style={styles.container}>
      <Text style={styles.header}>twimo</Text>
      <TextInput style={styles.input} autoCompleteType="off" />
      <TextInput
        style={styles.input}
        secureTextEntry={true}
        autoCompleteType="off"
      />
      <TouchableOpacity
        style={styles.buttonPrimary}
        onPress={() => navigation.navigate("List")}
      >
        <Text style={styles.buttonPrimaryText}>login</Text>
      </TouchableOpacity>
      <Button
        title="Sign up"
        onPress={() => console.log("Sir, we need a fucking signup page.")}
      ></Button>
    </View>
  );
};

// Define the styles
const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: color.white,
    alignItems: "center",
    justifyContent: "center"
  },
  header: {
    fontSize: 50,
    fontWeight: "bold",
    color: color.brand.normal
  },
  input: {
    height: 50,
    width: "70%",
    marginVertical: 10,
    paddingHorizontal: 10,
    borderRadius: 12,
    backgroundColor: color.brand.light,
    color: color.brand.dark
  },
  buttonPrimary: {
    backgroundColor: color.brand.normal,
    height: 50,
    width: "35%",
    borderRadius: 12,
    justifyContent: "center",
    alignItems: "center",
    marginTop: 20
  },
  buttonPrimaryText: {
    color: color.white,
    textTransform: "uppercase",
    fontWeight: "bold",
    fontSize: 20
  }
});

// Export the above defined screen
export default HomeScreen;
