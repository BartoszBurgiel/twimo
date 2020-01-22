// Import core functionalities from react & react-native library
import React from "react";
import { View, FlatList, StyleSheet } from "react-native";

// Import custom components
import Card from "../components/Card";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";

// Define Component
const ListScreen = ({ navigation }) => {
  // Define dummy data
  const dummyData = [
    {
      key: "1",
      name: "Cafe ABC",
      location: "XampleVillage",
      rating: 5,
      pricing: 3,
      picture: "https://placekitten.com/180/180"
    },
    {
      key: "2",
      name: "Cafe DEF",
      location: "XampleTown",
      rating: 2,
      pricing: 5,
      picture: "https://placekitten.com/180/180"
    },
    {
      key: "3",
      name: "Cafe GHI",
      location: "XampleCity",
      rating: 3,
      pricing: 3,
      picture: "https://placekitten.com/180/180"
    },
    {
      key: "4",
      name: "Cafe JKL",
      location: "XampleHills",
      rating: 4,
      pricing: 1,
      picture: "https://placekitten.com/180/180"
    },
    {
      key: "5",
      name: "Cafe XYZ",
      location: "XampleValley",
      rating: 3,
      pricing: 2,
      picture: "https://placekitten.com/180/180"
    }
  ];

  // Iterate iver the given data and display them as Cards in a FlatList
  return (
    <View style={styles.container}>
      <FlatList
        data={dummyData}
        showsVerticalScrollIndicator={false}
        renderItem={({ item }) => <Card data={item} navigation={navigation} />}
      />
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
  }
});

// Export the above defined component
export default ListScreen;
