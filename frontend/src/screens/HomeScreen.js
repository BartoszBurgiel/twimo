// Import core functionalities from react & react-native library
import React from "react";
import {
  View,
  Text,
  Image,
  Button,
  TextInput,
  TouchableOpacity,
  StyleSheet
} from "react-native";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";
import { Colors } from "react-native/Libraries/NewAppScreen";

//TODO: Implement real credential check!

// Define HomeScreen
const HomeScreen = ({ navigation }) => {
  return (
    <View style={styles.container}>
      <Image
        style={styles.wordmark}
        source={require("../../assets/wordmark.png")}
      />
      <TextInput
        style={styles.input}
        autoCompleteType="off"
        placeholder="Username"
        placeholderTextColor={color.brand.dark}
      />
      <TextInput
        style={styles.input}
        secureTextEntry={true}
        autoCompleteType="off"
        placeholder="Password"
        placeholderTextColor={color.brand.dark}
      />
      <TouchableOpacity
        style={styles.buttonPrimary}
        onPress={() => navigation.navigate("List")}
      >
        <Text style={styles.buttonPrimaryText}>Login</Text>
      </TouchableOpacity>
      <Button
        title="Sign up"
        onPress={() => navigation.navigate("SignUp")}
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
  wordmark: {
    marginVertical: 50
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
    color: color.black
  },
  buttonPrimary: {
    backgroundColor: color.brand.normal,
    height: 50,
    width: "30%",
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
