// Import core functionalities from react & react-native library
import React, { useState, useEffect } from "react";
import {
  View,
  StyleSheet,
  TextInput,
  Text,
  TouchableOpacity
} from "react-native";

// Import color utility
import color from "../../utils/colorPallet";

const SignUpScreen = ({ navigation }) => {
  // Render component
  return (
    <View style={styles.container}>
      <Text style={styles.header}>Sign up</Text>
      <Text>Name</Text>
      <TextInput
        style={styles.input}
        autoCompleteType="off"
        placeholder="e.g. Luke Skywalker"
      />
      <Text>Email</Text>
      <TextInput
        style={styles.input}
        autoCompleteType="off"
        placeholder="e.g. luke@jedi-mail.com"
      />
      <Text>Password</Text>
      <TextInput
        style={styles.input}
        autoCompleteType="off"
        placeholder="e.g. uSmellLikeChewbacca247"
      />
      <TouchableOpacity
        style={styles.buttonPrimary}
        onPress={() => navigation.navigate("List")}
      >
        <Text style={styles.buttonPrimaryText}>Let's go!</Text>
      </TouchableOpacity>
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
    fontSize: 30,
    fontWeight: "bold",
    marginBottom: 20,
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
    width: "50%",
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

// Export Screen
export default SignUpScreen;
