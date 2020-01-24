// Import core functionalities from react & react-native library
import React from "react";
import { View, FlatList, StyleSheet } from "react-native";

// Import custom components
import LocationCard from "../components/LocationCard";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";

// Define some dummy data
const dummyData = [
  {
    key: "1",
    name: "Starbucks Montabaur",
    location: "Montabaur",
    rating: 4,
    pricing: 3,
    feature: {
      coffee: true,
      wifi: true,
      power: true,
      music: false,
      food: true
    },
    picture:
      "https://www.heikaus-asset.com/wp-content/uploads/2018/01/Test_Starbucks-1024x731.jpg"
  },
  {
    key: "2",
    name: "Bäckerei Schäfer",
    location: "Westerburg",
    rating: 2,
    pricing: 5,
    feature: {
      coffee: true,
      wifi: true,
      power: false,
      music: true,
      food: true
    },
    picture:
      "https://www.schaefer-dein-baecker.de/sites/default/files/styles/flexslider_full/public/obu2_fcc7f0729f.jpg?itok=PBpOf_EO"
  },
  {
    key: "3",
    name: "Big Bender",
    location: "Herborn",
    rating: 3,
    pricing: 3,
    feature: {
      coffee: true,
      wifi: true,
      power: true,
      music: true,
      food: true
    },
    picture: "https://10619-2.s.cdn12.com/rests/original/337_507226641.jpg"
  },
  {
    key: "4",
    name: "Mühlenbäcker",
    location: "Rennerod",
    rating: 2,
    pricing: 2,
    feature: {
      coffee: true,
      wifi: true,
      power: false,
      music: false,
      food: true
    },
    picture:
      "https://www.die-muehlenbaecker.de/wp-content/uploads/IMG_5944-1024x683.jpg"
  },
  {
    key: "5",
    name: "Coffee Fellows",
    location: "Frankfurt",
    rating: 3,
    pricing: 2,
    feature: {
      coffee: true,
      wifi: true,
      power: true,
      music: false,
      food: true
    },
    picture:
      "https://www.coffee-fellows.com/wp-content/uploads/2017/03/coffee-fellows_cf_rs_breisgau_0017.jpg"
  }
];

// Define the ListScreen
const ListScreen = ({ navigation }) => {
  // Iterate over the given nodes and display them as LocationCards in a FlatList
  return (
    <View style={styles.container}>
      <FlatList
        data={dummyData}
        showsVerticalScrollIndicator={false}
        renderItem={({ item }) => (
          <LocationCard data={item} navigation={navigation} />
        )}
      />
    </View>
  );
};

// Define the styles
const styles = StyleSheet.create({
  container: {
    flex: 1,
    width: "100%",
    backgroundColor: color.white,
    alignItems: "center",
    justifyContent: "center"
  }
});

// Export the above defined screen
export default ListScreen;
