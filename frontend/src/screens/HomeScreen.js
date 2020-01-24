// Import core functionalities from react & react-native library
import React from "react";
import { View, Text, Button, StyleSheet } from "react-native";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";

// Define HomeScreen
const HomeScreen = ({ navigation }) => {
  return (
    <View style={styles.container}>
      <Text style={styles.textStyle}>twimo</Text>
      <Button
        title="Locations"
        onPress={() => navigation.navigate("List")}
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
  textStyle: {
    fontSize: 50,
    fontWeight: "bold",
    color: color.brand.normal
  }
});

// Export the above defined screen
export default HomeScreen;
