// Import core functionality dependencies from react & react-native
import React from "react";
import {
  View,
  ScrollView,
  FlatList,
  Text,
  StyleSheet,
  Image
} from "react-native";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";

// Import custom built components
import RatingStars from "../components/RatingStars";
import FeatureIcons from "../components/FeatureIcon";
import ReviewCard from "../components/ReviewCard";

// Define some dummy Review data
const dummyReviews = [
  {
    key: "1",
    name: "Ford Prefect",
    avatar: "https://api.adorable.io/avatars/50/fordprefect.png",
    rating: 5,
    content:
      "Klasse Handwerksbäcker, gemütliches und stilvolles Café mit hervorragenden Backwaren und bestem Kaffee. Unbedingt den Cafe Latte probieren - ein echter Genuss an diesen kalten Winterertagen!"
  },
  {
    key: "2",
    name: "Zaphod Beeblebrox",
    avatar: "https://api.adorable.io/avatars/50/zaphodbeeblebrox.png",
    rating: 4,
    content:
      "Klasse Cafe, gemütliches und stilvolles Ambiente mit hervorragenden Backwaren und gutem Kaffee. Unbedingt den Cafe Latte probieren - ein echter Genuss an diesen kalten Winterertagen!"
  },
  {
    key: "3",
    name: "Arthur Weasley",
    avatar: "https://api.adorable.io/avatars/50/arthurweasley.png",
    rating: 3,
    content:
      "Naja, gemütliches isses schon irgendwie. Aber der Durchzug ist heftig. Unbedingt ne Jacke mitnehmen - sonst gibt's ne fette Erkältung! Ansonsten kann man da schon hingehen..."
  },
  {
    key: "4",
    name: "Bellatrix Lestrange",
    avatar: "https://api.adorable.io/avatars/50/bellatrixlestrange.png",
    rating: 1,
    content:
      "Also ich find des nicht gut. Die Backwaren sind ganz ok, aber wenn'de kein Englisch kannst, bist in dem Laden uffgeschmissen. Da bleib ich beim Bäcker im Dorf!"
  }
];

// Define LocationDetailsScreen
const LocationDetailsScreen = props => {
  //FIXME: Virtualized Lists Warning

  // Grab locationData from the props and store it in data var
  const data = props.navigation.state.params.locationData;
  console.log(data);

  return (
    <ScrollView style={styles.container}>
      <View style={styles.imgWrapper}>
        <Image
          style={styles.image}
          source={{
            uri: data.picture
          }}
        ></Image>
      </View>
      <View style={styles.contentWrapper}>
        <Text style={styles.categoryTag}>coffeeshop</Text>
        <Text style={styles.headerText}>{data.name}</Text>
        <RatingStars colorFill={color.brand.normal} rating={data.rating} />
        <FeatureIcons features={data.feature} />
        <View style={styles.horizontalRule} />
        <Text style={styles.secondaryHeaderText}>Reviews</Text>
      </View>
      <FlatList
        data={dummyReviews}
        style={styles.contentWrapper}
        showsVerticalScrollIndicator={false}
        renderItem={({ item }) => <ReviewCard reviewData={item} />}
      />
    </ScrollView>
  );
};

// Define the styles
const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: color.white
  },
  imgWrapper: {
    alignItems: "center",
    marginBottom: 20
  },
  image: {
    height: 275,
    width: 500
  },
  contentWrapper: {
    paddingHorizontal: 30,
    marginBottom: 25
  },
  categoryTag: {
    fontSize: 20,
    fontWeight: "300",
    textTransform: "uppercase"
  },
  headerText: {
    fontSize: 30,
    fontWeight: "700"
  },
  horizontalRule: {
    borderBottomColor: color.gray.dark,
    borderBottomWidth: 1,
    marginVertical: 20
  },
  secondaryHeaderText: {
    fontSize: 20,
    fontWeight: "500"
  }
});

// Export the above defined screen
export default LocationDetailsScreen;
