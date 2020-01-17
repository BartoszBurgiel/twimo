// Import core functionalities from react & react-native library
import React from "react";
import { View, Text, FlatList, StyleSheet } from "react-native";

// Import custom components
import Card from "../components/Card";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";

// Define Component
const ListScreen = () => {
  // Define dummy data
  const dummyData = [
    {
      name: "Cafe ABC",
      location: "XampleVillage",
      rating: 5,
      pricing: 3,
      picture: "https://placekitten.com/180/180"
    },
    {
      name: "Cafe DEF",
      location: "XampleTown",
      rating: 2,
      pricing: 5,
      picture: "https://placekitten.com/180/180"
    },
    {
      name: "Cafe GHI",
      location: "XampleCity",
      rating: 3,
      pricing: 3,
      picture: "https://placekitten.com/180/180"
    },
    {
      name: "Cafe JKL",
      location: "XampleHills",
      rating: 4,
      pricing: 1,
      picture: "https://placekitten.com/180/180"
    },
    {
      name: "Cafe XYZ",
      location: "XampleValley",
      rating: 3,
      pricing: 2,
      picture: "https://placekitten.com/180/180"
    }
  ];

  // FIXME: List isn't scrollable yet
  return (
    <View style={styles.container}>
      <Card
        name={dummyData[0].name}
        location={dummyData[0].location}
        rating={dummyData[0].rating}
        pricing={dummyData[0].pricing}
      />
      <Card
        name={dummyData[1].name}
        location={dummyData[1].location}
        rating={dummyData[1].rating}
        pricing={dummyData[1].pricing}
      />
      <Card
        name={dummyData[2].name}
        location={dummyData[2].location}
        rating={dummyData[2].rating}
        pricing={dummyData[2].pricing}
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
